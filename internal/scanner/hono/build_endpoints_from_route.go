//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 단일 routeInfo에서 prefix 적용 + Zod 스키마 해석하여 Endpoint를 생성한다
package hono

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildEndpointsFromRoute(r routeInfo, vars map[string]bool, ctx *scanContext, fi *fileInfo, relPath string) []scanner.Endpoint {
	methods := expandAllMethod(r.Method)
	prefix := resolveRouteOwnerPrefix(r, ctx)
	fullPath := joinHonoPath(prefix, r.Path)
	oaPath := honoPathToOpenAPI(fullPath)
	pathParams := extractPathParams(fullPath)

	var endpoints []scanner.Endpoint
	for _, method := range methods {
		ep := buildOneEndpoint(method, oaPath, r, relPath, pathParams, ctx, fi)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}
