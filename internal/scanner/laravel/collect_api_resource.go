//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what Route::apiResource() 호출에서 CRUD 5개 라우트를 자동 생성한다
package laravel

// collectAPIResource extracts Route::apiResource('name', Controller::class) calls
// and expands them into 5 CRUD routes.
func collectAPIResource(fi fileInfo, prefix string, middleware []string) []routeInfo {
	var routes []routeInfo
	for _, call := range findAllByType(fi.root, "scoped_call_expression") {
		// Group-nested apiResource calls are handled by the recursive group
		// walk; skipping them here prevents duplicate, context-less routes.
		if isInsideGroupClosure(call, fi.root, fi) {
			continue
		}
		routes = append(routes, expandAPIResource(call, fi, prefix, middleware)...)
	}
	return routes
}
