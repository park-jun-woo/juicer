//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 파일에서 app.get(), app.post() 등 라우트 호출을 추출한다
package hono

func collectRoutes(fi *fileInfo, honoVars map[string]bool) []routeInfo {
	var routes []routeInfo
	calls := findAllByType(fi.Root, "call_expression")
	for _, call := range calls {
		ri := extractOneRoute(call, fi.Src, honoVars)
		if ri != nil {
			routes = append(routes, *ri)
		}
	}
	return routes
}
