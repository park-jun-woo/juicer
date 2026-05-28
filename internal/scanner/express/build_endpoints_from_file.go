//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 파일의 라우트를 추출하여 Endpoint 슬라이스를 생성한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildEndpointsFromFile(fi *fileInfo, routers map[string]bool, prefix, relPath string) []scanner.Endpoint {
	routes := extractRoutes(fi, routers)
	var endpoints []scanner.Endpoint
	for _, r := range routes {
		eps := buildEndpointsFromRoute(r, prefix, relPath)
		endpoints = append(endpoints, eps...)
	}
	return endpoints
}
