//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what 단일 Endpoint를 생성한다
package hono

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildOneEndpoint(method, oaPath string, r routeInfo, relPath string, pathParams []string, ctx *scanContext, fi *fileInfo) scanner.Endpoint {
	ep := scanner.Endpoint{
		Method:     method,
		Path:       oaPath,
		Handler:    r.Handler,
		File:       relPath,
		Line:       r.Line,
		Middleware: r.Middleware,
	}
	req := buildRequest(r, pathParams, ctx, fi)
	if req != nil {
		ep.Request = req
	}
	return ep
}
