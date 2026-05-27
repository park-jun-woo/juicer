//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what routeInfo로 scanner.Endpoint를 생성한다
package fastapi

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildEndpoint creates a scanner.Endpoint from a routeInfo.
func buildEndpoint(ri routeInfo) scanner.Endpoint {
	ep := scanner.Endpoint{
		Method:     ri.method,
		Path:       ri.path,
		Handler:    ri.handler,
		File:       ri.file,
		Line:       ri.line,
		Middleware: ri.middleware,
	}

	req := buildRequest(ri)
	if req != nil {
		ep.Request = req
	}

	respType := ri.responseModel
	if respType == "" {
		respType = ri.returnType
	}
	if ri.statusCode > 0 || respType != "" {
		ep.Responses = []scanner.Response{buildResponse(ri, respType)}
	}
	return ep
}
