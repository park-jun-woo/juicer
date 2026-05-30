//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what findRoute 테스트 헬퍼
package laravel

func findRoute(routes []routeInfo, path string) (routeInfo, bool) {
	for _, r := range routes {
		if r.path == path {
			return r, true
		}
	}
	return routeInfo{}, false
}
