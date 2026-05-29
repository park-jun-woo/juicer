//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 빌더 라우트 목록에서 중복(method+path+handler)을 제거한다
package actix

func deduplicateBuilderRoutes(routes []builderRoute) []builderRoute {
	seen := map[string]bool{}
	var result []builderRoute
	for _, r := range routes {
		key := r.method + " " + r.path + " " + r.handler
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, r)
	}
	return result
}
