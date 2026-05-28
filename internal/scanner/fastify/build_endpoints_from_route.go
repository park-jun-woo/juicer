//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 단일 routeInfo에서 Endpoint 슬라이스를 생성한다
package fastify

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildEndpointsFromRoute(r routeInfo, prefix, relPath string, src []byte) []scanner.Endpoint {
	methods := expandAllMethod(r.Method)
	fullPath := joinFastifyPath(prefix, r.Path)
	oaPath := fastifyPathToOpenAPI(fullPath)
	pathParams := extractPathParams(fullPath)
	var endpoints []scanner.Endpoint
	for _, method := range methods {
		ep := buildOneEndpoint(method, oaPath, r, relPath, pathParams, src)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}
