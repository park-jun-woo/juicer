//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what APIView 클래스에서 엔드포인트를 생성한다
package django

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildAPIViewEndpoints builds endpoints for an APIView class.
func buildAPIViewEndpoints(entry urlEntry, av *apiviewInfo, serializers map[string]serializerInfo) []scanner.Endpoint {
	openAPIPath := ensureLeadingSlash(djangoURLToOpenAPI(entry.pattern))
	urlParams := extractURLParams(entry.pattern)

	var endpoints []scanner.Endpoint
	for _, method := range av.methods {
		ep := scanner.Endpoint{
			Method:  method,
			Path:    openAPIPath,
			Handler: av.name + "." + strings.ToLower(method),
			File:    av.file,
			Line:    av.line,
		}
		addPathParams(&ep, urlParams)
		addWriteMethodBody(&ep, method, av.serializerClass, serializers)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}
