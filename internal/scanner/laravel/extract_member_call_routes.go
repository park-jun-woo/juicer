//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what member_call 하나를 그룹(->group) 또는 체인 단일 라우트(->get 등)로 분기 처리한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractMemberCallRoutes dispatches a member_call_expression to group handling
// (Route::prefix()->group()) or chained single-route handling
// (Route::middleware()->get()). Calls nested inside a group closure are deferred
// to the recursive group walk to keep the correct prefix/middleware context.
func extractMemberCallRoutes(mc *sitter.Node, fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	if isInsideGroupClosure(mc, fi.root, fi) {
		return nil
	}
	if groups := extractOneGroup(mc, fi, outerPrefix, outerMiddleware); groups != nil {
		return groups
	}
	if ri := extractChainedRoute(mc, fi, outerPrefix, outerMiddleware); ri != nil {
		return []routeInfo{*ri}
	}
	return nil
}
