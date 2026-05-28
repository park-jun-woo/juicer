//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what Actix-web 프로젝트를 스캔하여 엔드포인트를 추출한다
package actix

import (
	"fmt"
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func Scan(root string) (*scanner.ScanResult, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}
	rsFiles, err := findRsFiles(absRoot)
	if err != nil {
		return nil, fmt.Errorf("finding rs files: %w", err)
	}
	if len(rsFiles) == 0 {
		return &scanner.ScanResult{}, nil
	}

	files := parseAllFiles(absRoot, rsFiles)
	if len(files) == 0 {
		return &scanner.ScanResult{}, nil
	}

	sIdx := buildStructIndex(files)
	fieldCache := make(map[string][]scanner.Field)

	// Pass 1: Extract macro routes (#[get("/path")], #[post("/path")], etc.)
	var endpoints []scanner.Endpoint
	handlerFuncs := make(map[string]*handlerInfo)

	for _, fi := range files {
		macroRoutes := extractMacroRoutes(fi)
		for _, mr := range macroRoutes {
			ep := scanner.Endpoint{
				Method:  mr.method,
				Path:    mr.path,
				Handler: mr.handler,
				File:    fi.relPath,
			}

			// Extract path params from path (before extractors so types can be refined)
			pathParams := extractPathParams(mr.path)
			if len(pathParams) > 0 {
				ensureRequest(&ep)
				ep.Request.PathParams = pathParams
			}

			// Extract extractors from function signature
			if mr.funcNode != nil {
				exts := extractExtractors(mr.funcNode, fi.src)
				applyExtractors(&ep, exts, sIdx, fieldCache)

				// Extract responses
				responses := extractResponses(mr.funcNode, fi.src)
				if len(responses) > 0 {
					ep.Responses = responses
				}
			}

			endpoints = append(endpoints, ep)
		}

		// Collect handler functions for builder route resolution
		collectHandlerFuncs(fi, handlerFuncs)
	}

	// Pass 2: Extract builder routes (web::resource().route())
	for _, fi := range files {
		builderRoutes := extractBuilderRoutes(fi)
		for _, br := range builderRoutes {
			ep := scanner.Endpoint{
				Method:  br.method,
				Path:    br.path,
				Handler: br.handler,
			}

			// Extract path params before extractors
			pathParams := extractPathParams(br.path)
			if len(pathParams) > 0 {
				ensureRequest(&ep)
				ep.Request.PathParams = pathParams
			}

			// Resolve handler function for extractors
			if hi, ok := handlerFuncs[br.handler]; ok {
				ep.File = hi.file.relPath
				exts := extractExtractors(hi.funcNode, hi.file.src)
				applyExtractors(&ep, exts, sIdx, fieldCache)
				responses := extractResponses(hi.funcNode, hi.file.src)
				if len(responses) > 0 {
					ep.Responses = responses
				}
			}

			endpoints = append(endpoints, ep)
		}
	}

	// Pass 3: Apply scope prefixes to macro routes via .service(handler_name)
	for _, fi := range files {
		scopes := extractScopes(fi)
		for _, scope := range scopes {
			for _, handlerName := range scope.handlers {
				for i := range endpoints {
					if endpoints[i].Handler == handlerName && !hasPrefix(endpoints[i].Path, scope.prefix) {
						endpoints[i].Path = joinPath(scope.prefix, endpoints[i].Path)
					}
				}
			}
		}
	}

	return &scanner.ScanResult{Endpoints: endpoints}, nil
}

func applyExtractors(ep *scanner.Endpoint, exts []extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	for _, ext := range exts {
		switch ext.kind {
		case "path":
			applyPathExtractor(ep, ext, sIdx, cache)
		case "json":
			applyJSONExtractor(ep, ext, sIdx, cache)
		case "query":
			applyQueryExtractor(ep, ext, sIdx, cache)
		case "form":
			applyFormExtractor(ep, ext, sIdx, cache)
		}
	}
}

func applyPathExtractor(ep *scanner.Endpoint, ext extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	oaType := rustTypeToOpenAPI(ext.typeName)
	if oaType.Type != "object" {
		// Primitive type -- single path param; type already set from path
		ensureRequest(ep)
		for i := range ep.Request.PathParams {
			ep.Request.PathParams[i].Type = oaType.Type
		}
		return
	}
	// Struct type -- expand fields as path params
	fields := resolveStructFields(ext.typeName, sIdx, cache)
	if len(fields) > 0 {
		ensureRequest(ep)
		var params []scanner.Param
		for _, f := range fields {
			params = append(params, scanner.Param{Name: f.JSON, Type: f.Type})
		}
		ep.Request.PathParams = params
	}
}

func applyJSONExtractor(ep *scanner.Endpoint, ext extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	ensureRequest(ep)
	ep.Request.Body = &scanner.Body{
		VarName:  "body",
		Method:   "json",
		TypeName: ext.typeName,
	}
	fields := resolveStructFields(ext.typeName, sIdx, cache)
	if len(fields) > 0 {
		ep.Request.Body.Fields = fields
	}
}

func applyQueryExtractor(ep *scanner.Endpoint, ext extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	fields := resolveStructFields(ext.typeName, sIdx, cache)
	if len(fields) == 0 {
		return
	}
	ensureRequest(ep)
	for _, f := range fields {
		ep.Request.Query = append(ep.Request.Query, scanner.Param{
			Name: f.JSON,
			Type: f.Type,
		})
	}
}

func applyFormExtractor(ep *scanner.Endpoint, ext extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	fields := resolveStructFields(ext.typeName, sIdx, cache)
	if len(fields) == 0 {
		return
	}
	ensureRequest(ep)
	for _, f := range fields {
		ep.Request.FormFields = append(ep.Request.FormFields, scanner.Param{
			Name: f.JSON,
			Type: f.Type,
		})
	}
}

func ensureRequest(ep *scanner.Endpoint) {
	if ep.Request == nil {
		ep.Request = &scanner.Request{}
	}
}

func hasPrefix(path, prefix string) bool {
	if prefix == "" || prefix == "/" {
		return true
	}
	return len(path) >= len(prefix) && path[:len(prefix)] == prefix
}

type handlerInfo struct {
	funcNode *sitter.Node
	file     *fileInfo
}
