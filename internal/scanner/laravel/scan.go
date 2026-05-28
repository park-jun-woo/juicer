//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Laravel 프로젝트를 3-pass로 스캔하여 엔드포인트를 추출한다
package laravel

import (
	"fmt"
	"path/filepath"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// Scan scans a Laravel project root and extracts endpoints.
// Pass 1: collect route structure from routes/api.php and routes/web.php.
// Pass 2: resolve controller methods — extract parameter types and return info.
// Pass 3: resolve FormRequest rules and Resource response fields.
func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}

	phpFiles, err := findPHPFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding php files: %w", err)
	}
	if len(phpFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	// Parse all PHP files
	parsedFiles := make(map[string]*fileInfo)
	for _, f := range phpFiles {
		fi, err := parseFile(absRoot, f)
		if err != nil {
			continue
		}
		parsedFiles[fi.relPath] = fi
	}
	if len(parsedFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	// Pass 1: collect routes from route files
	routes := collectAllRoutes(parsedFiles)
	if len(routes) == 0 {
		return &scanner.ScanResult{}, nil
	}

	// Pass 2 + 3: build endpoints with controller and FormRequest resolution
	endpoints := buildEndpoints(absRoot, routes, parsedFiles)

	return &scanner.ScanResult{Endpoints: endpoints}, nil
}

// collectAllRoutes gathers routes from route files (routes/api.php, routes/web.php).
func collectAllRoutes(parsedFiles map[string]*fileInfo) []routeInfo {
	var routes []routeInfo
	routeFiles := []string{
		"routes/api.php",
		"routes/web.php",
	}
	for _, rf := range routeFiles {
		fi, ok := parsedFiles[rf]
		if !ok {
			continue
		}
		// Determine default prefix for api.php
		prefix := ""
		if strings.HasSuffix(rf, "api.php") {
			prefix = "api"
		}
		routes = append(routes, collectRoutes(*fi, prefix, nil)...)
		routes = append(routes, collectAPIResource(*fi, prefix, nil)...)
		routes = append(routes, extractRouteGroups(*fi, prefix, nil)...)
	}
	return routes
}

// buildEndpoints converts raw route info into scanner.Endpoint with resolved
// controller params and FormRequest fields.
func buildEndpoints(absRoot string, routes []routeInfo, parsedFiles map[string]*fileInfo) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, ri := range routes {
		ep := scanner.Endpoint{
			Method:     ri.method,
			Path:       ri.path,
			Handler:    buildHandlerName(ri.controller, ri.action),
			File:       ri.file,
			Line:       ri.line,
			Middleware: ri.middleware,
		}

		// Extract path params from URL
		pathParams := extractURLParams(ri.path)

		// Pass 2: Resolve controller method
		var cm *controllerMethod
		if ri.controller != "" && ri.action != "" {
			controllerFI := resolveController(absRoot, ri.controller, parsedFiles)
			if controllerFI != nil {
				cm = extractControllerMethod(controllerFI, ri.action)
			}
		}

		// Apply controller type hints to path params
		pathParams = applyControllerParamTypes(pathParams, cm)

		// Set up request
		if len(pathParams) > 0 || (cm != nil && cm.formRequestRef != "") {
			req := &scanner.Request{
				PathParams: pathParams,
			}

			// Pass 3: Resolve FormRequest
			if cm != nil && cm.formRequestRef != "" {
				fields := extractFormRequest(absRoot, cm.formRequestRef, parsedFiles)
				if len(fields) > 0 {
					req.Body = &scanner.Body{
						VarName:  "request",
						Method:   "json",
						TypeName: cm.formRequestRef,
						Fields:   fields,
					}
				}
			}
			ep.Request = req
		}

		// Pass 3: Resolve response Resource
		if cm != nil {
			responses := extractResponsesFromMethod(absRoot, cm, parsedFiles)
			if len(responses) > 0 {
				ep.Responses = responses
			}
		}

		endpoints = append(endpoints, ep)
	}
	return endpoints
}

// buildHandlerName creates handler string from controller and action.
func buildHandlerName(controller, action string) string {
	if controller == "" && action == "" {
		return "closure"
	}
	if controller == "" {
		return action
	}
	if action == "" {
		return controller
	}
	return controller + "@" + action
}

// extractResponsesFromMethod extracts response info from controller method return statements.
func extractResponsesFromMethod(absRoot string, cm *controllerMethod, parsedFiles map[string]*fileInfo) []scanner.Response {
	var responses []scanner.Response
	for _, retNode := range cm.returnNodes {
		resp := extractOneResponse(absRoot, retNode, cm.src, parsedFiles)
		if resp != nil {
			responses = append(responses, *resp)
		}
	}
	return responses
}

// extractOneResponse extracts response info from a single return statement.
func extractOneResponse(absRoot string, retNode *sitter.Node, src []byte, parsedFiles map[string]*fileInfo) *scanner.Response {
	text := nodeText(retNode, src)

	// new SomeResource($var) -> JSON resource
	objCreations := findAllByType(retNode, "object_creation_expression")
	for _, oc := range objCreations {
		nameNode := findChildByType(oc, "name")
		if nameNode != nil {
			resName := nodeText(nameNode, src)
			if strings.HasSuffix(resName, "Resource") || strings.HasSuffix(resName, "Collection") {
				fields := extractResourceFields(absRoot, resName, parsedFiles)
				return &scanner.Response{
					Status:   "200",
					Kind:     "json",
					TypeName: resName,
					Fields:   fields,
				}
			}
		}
	}

	// SomeResource::collection($var) -> JSON array response
	scopedCalls := findAllByType(retNode, "scoped_call_expression")
	for _, sc := range scopedCalls {
		nameNode := findChildByType(sc, "name")
		if nameNode == nil {
			continue
		}
		name := nodeText(nameNode, src)
		// Check for ::collection() pattern
		secondName := ""
		foundFirst := false
		for i := 0; i < int(sc.ChildCount()); i++ {
			child := sc.Child(i)
			if child.Type() == "name" {
				if !foundFirst {
					foundFirst = true
					continue
				}
				secondName = nodeText(child, src)
				break
			}
		}
		if secondName == "collection" && (strings.HasSuffix(name, "Resource") || strings.HasSuffix(name, "Collection")) {
			fields := extractResourceFields(absRoot, name, parsedFiles)
			return &scanner.Response{
				Status:   "200",
				Kind:     "json",
				TypeName: "[]" + name,
				Fields:   fields,
			}
		}
	}

	// response()->json($data, 201) or response()->json(...)
	if strings.Contains(text, "response()->json(") {
		status := "200"
		// Try to extract status code
		memberCalls := findAllByType(retNode, "member_call_expression")
		for _, mc := range memberCalls {
			mName := ""
			for i := 0; i < int(mc.ChildCount()); i++ {
				child := mc.Child(i)
				if child.Type() == "name" {
					mName = nodeText(child, src)
				}
			}
			if mName == "json" {
				jsonArgs := findChildByType(mc, "arguments")
				if jsonArgs != nil {
					argList := childrenOfType(jsonArgs, "argument")
					if len(argList) >= 2 {
						// Second arg is status code
						code := nodeText(argList[1], src)
						code = strings.TrimSpace(code)
						if code != "" {
							status = code
						}
					}
				}
			}
		}
		return &scanner.Response{
			Status: status,
			Kind:   "json",
		}
	}

	// response()->noContent()
	if strings.Contains(text, "noContent()") {
		return &scanner.Response{
			Status: "204",
			Kind:   "empty",
		}
	}

	return nil
}
