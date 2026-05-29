//ff:func feature=scan type=extract control=sequence topic=django
//ff:what DRF가 아닌 순수 Django 뷰(함수형/CBV)를 GET 엔드포인트 1건으로 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildPlainViewEndpoints is a fallback for plain Django views (function or CBV)
// that are not DRF ViewSet/APIView/@api_view. It emits a single GET endpoint.
// HTTP method dispatch inside the view body is not statically resolvable, so GET
// is used as a documented default (known limitation).
func buildPlainViewEndpoints(entry urlEntry, viewName string) []scanner.Endpoint {
	openAPIPath := ensureLeadingSlash(djangoURLToOpenAPI(entry.pattern))
	ep := scanner.Endpoint{
		Method:  "GET",
		Path:    openAPIPath,
		Handler: viewName,
	}
	addPathParams(&ep, extractURLParams(entry.pattern))
	return []scanner.Endpoint{ep}
}
