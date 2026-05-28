//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 라우트에서 request 정보를 구성한다 (path params + schema 기반)
package fastify

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildRequest(r routeInfo, pathParams []string, src []byte) (scanner.Request, bool) {
	var req scanner.Request
	hasRequest := false
	if len(pathParams) > 0 {
		req.PathParams = buildPathParams(pathParams)
		hasRequest = true
	}
	if r.Schema == nil {
		return req, hasRequest
	}
	si := extractJSONSchema(r.Schema, src)
	if si == nil {
		return req, hasRequest
	}
	applySchemaToRequest(si, src, &req, &hasRequest)
	return req, hasRequest
}
