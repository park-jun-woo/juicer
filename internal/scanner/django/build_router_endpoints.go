//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what Router 등록된 ViewSet에서 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildRouterEndpoints builds endpoints for all router-registered ViewSets.
func buildRouterEndpoints(regs []routerRegistration, viewsets []viewsetInfo, serializers map[string]serializerInfo) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, reg := range regs {
		vsName := resolveViewName(reg.viewsetName)
		vs := findViewSet(viewsets, vsName)
		if vs == nil {
			continue
		}
		endpoints = append(endpoints, buildViewSetEndpoints(reg, vs, serializers)...)
	}
	return endpoints
}
