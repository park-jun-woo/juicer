//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 routeInfo에서 (all 확장 포함) Endpoint를 생성한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildEndpointsFromRoute(r routeInfo, prefix, relPath string) []scanner.Endpoint {
	methods := expandAllMethod(r.Method)
	fullPath := joinExpressPath(prefix, r.Path)
	oaPath := expressPathToOpenAPI(fullPath)
	pathParams := extractPathParams(fullPath)
	var endpoints []scanner.Endpoint
	for _, method := range methods {
		ep := buildOneEndpoint(method, oaPath, r, relPath, pathParams)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}
