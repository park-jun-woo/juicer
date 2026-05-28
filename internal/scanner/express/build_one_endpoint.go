//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 단일 Endpoint를 생성한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildOneEndpoint(method, oaPath string, r routeInfo, relPath string, pathParams []string, ctx *scanContext, fi *fileInfo) scanner.Endpoint {
	ep := scanner.Endpoint{
		Method:     method,
		Path:       oaPath,
		Handler:    r.Handler,
		File:       relPath,
		Line:       r.Line,
		Middleware: r.Middleware,
		AuthLevel:  r.AuthLevel,
		Roles:      r.Roles,
	}
	req := buildRequest(r, pathParams, ctx, fi)
	if req != nil {
		ep.Request = req
	}
	ep.Responses = extractResponses(fi, r)
	return ep
}
