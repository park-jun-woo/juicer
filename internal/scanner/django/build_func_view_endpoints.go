//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 함수 기반 뷰에서 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildFuncViewEndpoints builds endpoints for a function-based view.
func buildFuncViewEndpoints(entry urlEntry, fv *funcViewInfo) []scanner.Endpoint {
	openAPIPath := ensureLeadingSlash(djangoURLToOpenAPI(entry.pattern))
	urlParams := extractURLParams(entry.pattern)

	var endpoints []scanner.Endpoint
	for _, method := range fv.methods {
		ep := scanner.Endpoint{
			Method:  method,
			Path:    openAPIPath,
			Handler: fv.name,
			File:    fv.file,
			Line:    fv.line,
		}
		addPathParams(&ep, urlParams)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}
