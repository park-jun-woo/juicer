//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what Route::group([배열], fn) 형태의 scoped 그룹 호출들을 수집한다
package laravel

// collectScopedGroups gathers array-form Route::group([...], fn) calls. Groups
// nested inside another group closure are deferred to the recursive group walk
// so that the outer prefix/middleware are applied exactly once.
func collectScopedGroups(fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	var routes []routeInfo
	for _, call := range findAllByType(fi.root, "scoped_call_expression") {
		if isInsideGroupClosure(call, fi.root, fi) {
			continue
		}
		routes = append(routes, extractScopedGroup(call, fi, outerPrefix, outerMiddleware)...)
	}
	return routes
}
