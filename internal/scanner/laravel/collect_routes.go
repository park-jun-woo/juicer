//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what routes/api.php, routes/web.php에서 Route::get/post 등 개별 라우트를 수집한다
package laravel

// collectRoutes extracts Route::get/post/put/patch/delete calls from a file.
func collectRoutes(fi fileInfo, prefix string, middleware []string) []routeInfo {
	var routes []routeInfo
	for _, call := range findAllByType(fi.root, "scoped_call_expression") {
		// Routes nested inside a ->group(closure) are owned by
		// extractRouteGroups (which supplies the group prefix/middleware).
		// Skipping them here avoids duplicate, context-less endpoints.
		if isInsideGroupClosure(call, fi.root, fi) {
			continue
		}
		ri := extractOneRoute(call, fi, prefix, middleware)
		if ri != nil {
			routes = append(routes, *ri)
		}
	}
	return routes
}
