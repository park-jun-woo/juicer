//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 데코레이터에서 HTTP 라우트를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractRoutes finds all decorated route functions in the AST and returns routeInfo slices.
// prefixes maps router variable names to their resolved path prefixes.
func extractRoutes(root *sitter.Node, src []byte, prefixes map[string]string, file string) []routeInfo {
	var routes []routeInfo
	defs := findAllByType(root, "decorated_definition")
	for _, def := range defs {
		ri := extractOneRoute(def, src, prefixes, file)
		if ri != nil {
			routes = append(routes, *ri)
		}
	}
	return routes
}
