//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Route::method('/path', handler) 호출 하나를 routeInfo로 추출한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractOneRoute extracts a single Route::method('/path', handler) call.
func extractOneRoute(call *sitter.Node, fi fileInfo, prefix string, middleware []string) *routeInfo {
	if call.ChildCount() < 4 {
		return nil
	}
	scope := findChildByType(call, "name")
	if scope == nil || nodeText(scope, fi.src) != "Route" {
		return nil
	}

	methodName := secondScopedName(call, fi.src)
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
	pathStr := extractStringContent(argNodes[0], fi.src)
	if pathStr == "" {
		return nil
	}

	controller, action := extractControllerAction(argNodes[1], fi.src)
	return &routeInfo{
		method:     upperMethod,
		path:       joinLaravelPath(prefix, pathStr),
		controller: controller,
		action:     action,
		file:       fi.relPath,
		line:       int(call.StartPoint().Row) + 1,
		middleware: copyMiddleware(middleware),
	}
}
