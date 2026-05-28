//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what routes/api.php, routes/web.php에서 Route::get/post 등 개별 라우트를 수집한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// routeInfo holds a single extracted route.
type routeInfo struct {
	method     string // HTTP method (uppercase)
	path       string // URL path
	controller string // controller class name
	action     string // controller method name
	file       string // source file relative path
	line       int    // source line number
	middleware []string
}

// collectRoutes extracts Route::get/post/put/patch/delete calls from a file.
func collectRoutes(fi fileInfo, prefix string, middleware []string) []routeInfo {
	var routes []routeInfo
	calls := findAllByType(fi.root, "scoped_call_expression")
	for _, call := range calls {
		ri := extractOneRoute(call, fi, prefix, middleware)
		if ri != nil {
			routes = append(routes, *ri)
		}
	}
	return routes
}

// extractOneRoute extracts a single Route::method('/path', handler) call.
func extractOneRoute(call *sitter.Node, fi fileInfo, prefix string, middleware []string) *routeInfo {
	// Check structure: Route::method(args)
	if call.ChildCount() < 4 {
		return nil
	}
	scope := findChildByType(call, "name")
	if scope == nil || nodeText(scope, fi.src) != "Route" {
		return nil
	}

	// Find the method name (second "name" child after "::")
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
	if methodName == "" {
		return nil
	}

	upperMethod, ok := httpMethods[strings.ToLower(methodName)]
	if !ok {
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

	// First argument: path string
	pathStr := extractStringContent(argNodes[0], fi.src)
	if pathStr == "" {
		return nil
	}

	// Build full path with prefix
	fullPath := joinLaravelPath(prefix, pathStr)

	// Second argument: handler [Controller::class, 'method'] or closure
	controller, action := extractControllerAction(argNodes[1], fi.src)

	mw := make([]string, len(middleware))
	copy(mw, middleware)

	return &routeInfo{
		method:     upperMethod,
		path:       fullPath,
		controller: controller,
		action:     action,
		file:       fi.relPath,
		line:       int(call.StartPoint().Row) + 1,
		middleware: mw,
	}
}

// extractStringContent extracts the string content from a string node.
func extractStringContent(node *sitter.Node, src []byte) string {
	strNodes := findAllByType(node, "string_content")
	if len(strNodes) > 0 {
		return nodeText(strNodes[0], src)
	}
	// Try encapsed_string / string
	strLit := findChildByType(node, "string")
	if strLit != nil {
		return unquotePHP(nodeText(strLit, src))
	}
	text := nodeText(node, src)
	return unquotePHP(text)
}

// extractControllerAction extracts controller class and method from
// [Controller::class, 'method'] array syntax.
func extractControllerAction(node *sitter.Node, src []byte) (string, string) {
	arr := findChildByType(node, "array_creation_expression")
	if arr == nil {
		// Check if the node itself is an array
		if node.Type() == "array_creation_expression" {
			arr = node
		} else {
			return "", ""
		}
	}
	elems := childrenOfType(arr, "array_element_initializer")
	if len(elems) < 2 {
		return "", ""
	}
	// First element: Controller::class
	classAccess := findChildByType(elems[0], "class_constant_access_expression")
	if classAccess == nil {
		classAccess = findAllByType(elems[0], "class_constant_access_expression")[0]
	}
	controller := ""
	if classAccess != nil {
		nameNode := findChildByType(classAccess, "name")
		if nameNode != nil {
			controller = nodeText(nameNode, src)
		}
	}
	// Second element: 'method'
	action := extractStringContent(elems[1], src)
	return controller, action
}

// joinLaravelPath joins prefix and path segments into a clean URL path.
func joinLaravelPath(prefix, path string) string {
	prefix = strings.TrimRight(prefix, "/")
	path = strings.TrimLeft(path, "/")
	if prefix == "" && path == "" {
		return "/"
	}
	if prefix == "" {
		return "/" + path
	}
	if path == "" {
		return "/" + prefix
	}
	return "/" + prefix + "/" + path
}
