//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 데코레이터에서 Flask HTTP 라우트를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// extractRoutes finds all decorated route functions in the AST and returns routeInfo slices.
func extractRoutes(root *sitter.Node, src []byte, bpPrefixes blueprintPrefix, file string) []routeInfo {
	var routes []routeInfo
	defs := findAllByType(root, "decorated_definition")
	for _, def := range defs {
		rs := extractOneRoute(def, src, bpPrefixes, file)
		routes = append(routes, rs...)
	}
	return routes
}
