//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what Route::apiResource() 호출에서 CRUD 5개 라우트를 자동 생성한다
package laravel

import (
	"fmt"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// collectAPIResource extracts Route::apiResource('name', Controller::class) calls
// and expands them into 5 CRUD routes.
func collectAPIResource(fi fileInfo, prefix string, middleware []string) []routeInfo {
	var routes []routeInfo
	calls := findAllByType(fi.root, "scoped_call_expression")
	for _, call := range calls {
		rs := expandAPIResource(call, fi, prefix, middleware)
		routes = append(routes, rs...)
	}
	return routes
}

// expandAPIResource expands a single Route::apiResource call into CRUD routes.
func expandAPIResource(call *sitter.Node, fi fileInfo, prefix string, middleware []string) []routeInfo {
	if call.ChildCount() < 4 {
		return nil
	}
	scope := findChildByType(call, "name")
	if scope == nil || nodeText(scope, fi.src) != "Route" {
		return nil
	}

	methodName := ""
	foundScope := false
	for i := 0; i < int(call.ChildCount()); i++ {
		child := call.Child(i)
		if child.Type() == "name" {
			if !foundScope {
				foundScope = true
				continue
			}
			methodName = nodeText(child, fi.src)
			break
		}
	}
	if methodName != "apiResource" {
		return nil
	}

	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	argNodes := childrenOfType(args, "argument")
	if len(argNodes) < 2 {
		return nil
	}

	// First arg: resource name (e.g., 'users' or 'users.posts')
	resourceName := extractStringContent(argNodes[0], fi.src)
	if resourceName == "" {
		return nil
	}

	// Second arg: Controller::class
	controller := extractClassReference(argNodes[1], fi.src)

	// Build resource path from dotted name (e.g., "users.posts" -> "/users/{user}/posts")
	basePath, paramName := buildResourcePath(resourceName)
	fullBase := joinLaravelPath(prefix, basePath)

	mw := make([]string, len(middleware))
	copy(mw, middleware)

	var routes []routeInfo
	for _, action := range apiResourceActions {
		suffix := action.suffix
		if action.hasParam {
			suffix = fmt.Sprintf(suffix, paramName)
		}
		routes = append(routes, routeInfo{
			method:     action.method,
			path:       fullBase + suffix,
			controller: controller,
			action:     action.action,
			file:       fi.relPath,
			line:       int(call.StartPoint().Row) + 1,
			middleware: mw,
		})
	}
	return routes
}

// extractClassReference extracts a class name from Controller::class expression.
func extractClassReference(node *sitter.Node, src []byte) string {
	classAccess := findChildByType(node, "class_constant_access_expression")
	if classAccess == nil {
		nodes := findAllByType(node, "class_constant_access_expression")
		if len(nodes) > 0 {
			classAccess = nodes[0]
		}
	}
	if classAccess == nil {
		return ""
	}
	nameNode := findChildByType(classAccess, "name")
	if nameNode == nil {
		return ""
	}
	return nodeText(nameNode, src)
}

// buildResourcePath converts a dotted resource name to URL path segments.
// "users" -> ("users", "user")
// "users.posts" -> ("users/{user}/posts", "post")
func buildResourcePath(name string) (string, string) {
	parts := strings.Split(name, ".")
	if len(parts) == 1 {
		return parts[0], singularize(parts[0])
	}
	var path string
	for i, part := range parts {
		if i > 0 {
			path += "/" + part
		} else {
			path = part
		}
		if i < len(parts)-1 {
			path += "/{" + singularize(part) + "}"
		}
	}
	return path, singularize(parts[len(parts)-1])
}

// singularize performs naive English singularization (remove trailing 's').
func singularize(s string) string {
	if strings.HasSuffix(s, "ies") {
		return s[:len(s)-3] + "y"
	}
	if strings.HasSuffix(s, "ses") || strings.HasSuffix(s, "xes") || strings.HasSuffix(s, "zes") {
		return s[:len(s)-2]
	}
	if strings.HasSuffix(s, "s") && !strings.HasSuffix(s, "ss") {
		return s[:len(s)-1]
	}
	return s
}
