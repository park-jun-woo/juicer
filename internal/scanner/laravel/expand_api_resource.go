//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what Route::apiResource 호출 하나를 CRUD 라우트들로 펼친다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// expandAPIResource expands a single Route::apiResource call into CRUD routes.
func expandAPIResource(call *sitter.Node, fi fileInfo, prefix string, middleware []string) []routeInfo {
	if call.ChildCount() < 4 {
		return nil
	}
	scope := findChildByType(call, "name")
	if scope == nil || nodeText(scope, fi.src) != "Route" {
		return nil
	}
	if secondScopedName(call, fi.src) != "apiResource" {
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
	resourceName := extractStringContent(argNodes[0], fi.src)
	if resourceName == "" {
		return nil
	}

	controller := extractClassReference(argNodes[1], fi.src)
	basePath, paramName := buildResourcePath(resourceName)
	fullBase := joinLaravelPath(prefix, basePath)
	mw := copyMiddleware(middleware)
	line := int(call.StartPoint().Row) + 1

	var routes []routeInfo
	for _, action := range apiResourceActions {
		routes = append(routes, apiResourceRoute(action, fullBase, paramName, controller, fi.relPath, line, mw))
	}
	return routes
}
