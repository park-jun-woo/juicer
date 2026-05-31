//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what path()로 매핑된 ViewSet에서 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildURLMappedViewSetEndpoints builds endpoints for a ViewSet mapped via path().
func buildURLMappedViewSetEndpoints(entry urlEntry, vs *viewsetInfo, serializers map[string]serializerInfo) []scanner.Endpoint {
	openAPIPath := ensureLeadingSlash(djangoURLToOpenAPI(entry.pattern))
	urlParams := extractURLParams(entry.pattern)
	methods := urlEntryViewSetMethods(entry, vs)

	var endpoints []scanner.Endpoint
	for _, am := range methods {
		ep := scanner.Endpoint{
			Method:  am.method,
			Path:    openAPIPath,
			Handler: vs.name + "." + am.action,
			File:    vs.file,
			Line:    vs.line,
		}
		addPathParams(&ep, urlParams)
		addSerializerInfo(&ep, am, vs.serializerClass, serializers)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}
