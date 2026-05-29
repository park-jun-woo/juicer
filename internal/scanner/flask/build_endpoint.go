//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what routeInfo를 scanner.Endpoint로 변환한다
package flask

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildEndpoint converts a routeInfo to a scanner.Endpoint.
func buildEndpoint(ri routeInfo) scanner.Endpoint {
	ep := scanner.Endpoint{
		Method:  ri.method,
		Path:    ri.path,
		Handler: ri.handler,
		File:    ri.file,
		Line:    ri.line,
	}

	pathParams := urlParamsToScannerParams(ri.params)
	if len(pathParams) > 0 {
		scanner.EnsureRequest(&ep)
		ep.Request.PathParams = pathParams
	}

	applyFormFields(&ep, ri.formFields)
	applyJSONBody(&ep, ri)

	return ep
}
