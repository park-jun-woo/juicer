//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 클로저 본문 안의 라우트/apiResource/중첩 그룹을 수집한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// collectRoutesFromBody extracts routes from inside a closure body.
func collectRoutesFromBody(body *sitter.Node, fi fileInfo, prefix string, middleware []string) []routeInfo {
	bodyFI := fileInfo{
		absPath: fi.absPath,
		relPath: fi.relPath,
		src:     fi.src,
		root:    body,
	}
	var routes []routeInfo
	routes = append(routes, collectRoutes(bodyFI, prefix, middleware)...)
	routes = append(routes, collectAPIResource(bodyFI, prefix, middleware)...)

	for _, mc := range findAllByType(body, "member_call_expression") {
		routes = append(routes, extractMemberCallRoutes(mc, bodyFI, prefix, middleware)...)
	}
	routes = append(routes, collectScopedGroups(bodyFI, prefix, middleware)...)
	return routes
}
