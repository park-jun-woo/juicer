//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what URL 패턴에서 View/ViewSet/함수 뷰 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildURLEntryEndpoints builds endpoints for URL pattern entries.
func buildURLEntryEndpoints(entries []urlEntry, viewsets []viewsetInfo, apiviews []apiviewInfo, funcViews []funcViewInfo, serializers map[string]serializerInfo) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, entry := range entries {
		if entry.isInclude {
			continue
		}
		eps := buildSingleURLEntryEndpoints(entry, viewsets, apiviews, funcViews, serializers)
		endpoints = append(endpoints, eps...)
	}
	return endpoints
}
