//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what Route::prefix()->group() / Route::middleware()->group() 체인에서 그룹 라우트를 추출한다
package laravel

// extractRouteGroups extracts routes from Route::prefix()->group(),
// Route::middleware()->group() chain calls, chained single routes such as
// Route::middleware()->get(), and array-form Route::group([...], fn) calls.
func extractRouteGroups(fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	var routes []routeInfo
	for _, mc := range findAllByType(fi.root, "member_call_expression") {
		routes = append(routes, extractMemberCallRoutes(mc, fi, outerPrefix, outerMiddleware)...)
	}
	routes = append(routes, collectScopedGroups(fi, outerPrefix, outerMiddleware)...)
	return routes
}
