//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 파일의 라우트를 추출하여 Endpoint 슬라이스를 생성한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildEndpointsFromFile(fi *fileInfo, routers map[string]bool, file, relPath string, ctx *scanContext) []scanner.Endpoint {
	routes := extractRoutes(fi, routers)
	var endpoints []scanner.Endpoint
	for _, r := range routes {
		// 라우트가 등록된 라우터 인스턴스의 prefix 목록을 적용한다.
		// 한 라우터가 여러 prefix에 마운트되면 각 prefix마다 endpoint를 생성한다.
		for _, prefix := range routePrefixes(ctx, file, r.Router) {
			eps := buildEndpointsFromRoute(r, prefix, relPath, ctx, fi)
			endpoints = append(endpoints, eps...)
		}
	}
	return endpoints
}
