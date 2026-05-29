//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 라우트 정보들을 컨트롤러/FormRequest 해석을 거쳐 엔드포인트로 변환한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildEndpoints converts raw route info into scanner.Endpoint with resolved
// controller params and FormRequest fields.
func buildEndpoints(absRoot string, routes []routeInfo, parsedFiles map[string]*fileInfo) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, ri := range routes {
		endpoints = append(endpoints, buildOneEndpoint(absRoot, ri, parsedFiles))
	}
	return endpoints
}
