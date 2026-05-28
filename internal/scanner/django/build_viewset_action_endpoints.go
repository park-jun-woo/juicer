//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what ViewSet의 @action 커스텀 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildViewSetActionEndpoints builds endpoints for a ViewSet @action method.
func buildViewSetActionEndpoints(prefix string, ai actionInfo, vs *viewsetInfo) []scanner.Endpoint {
	actionPath := ai.urlPath
	if actionPath == "" {
		actionPath = ai.name
	}
	var endpoints []scanner.Endpoint
	for _, method := range ai.methods {
		ep := buildOneActionEndpoint(prefix, actionPath, method, ai, vs)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}
