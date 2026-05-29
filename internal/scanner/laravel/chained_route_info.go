//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 체인 라우트 member_call의 (path, handler) 인자로 routeInfo를 만든다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// chainedRouteInfo builds a routeInfo from the terminal HTTP-verb member call
// of a chained route, using its (path, handler) arguments.
func chainedRouteInfo(mc *sitter.Node, fi fileInfo, method, prefix string, middleware []string) *routeInfo {
	args := findChildByType(mc, "arguments")
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
		method:     method,
		path:       joinLaravelPath(prefix, pathStr),
		controller: controller,
		action:     action,
		file:       fi.relPath,
		line:       int(mc.StartPoint().Row) + 1,
		middleware: copyMiddleware(middleware),
	}
}
