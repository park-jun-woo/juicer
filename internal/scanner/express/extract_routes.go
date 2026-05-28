//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 파일에서 모든 라우트(단일 + 체인)를 추출한다
package express

func extractRoutes(fi *fileInfo, routers map[string]bool) []routeInfo {
	var routes []routeInfo
	calls := findAllByType(fi.Root, "call_expression")
	processed := make(map[uintptr]bool)
	for _, call := range calls {
		id := uintptr(call.StartByte())
		if processed[id] {
			continue
		}
		ri := extractRouteFromCall(call, fi.Src, routers, processed)
		routes = append(routes, ri...)
	}
	return routes
}
