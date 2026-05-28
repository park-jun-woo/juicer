//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 단일 리소스의 엔드포인트를 생성한다
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildResourceEndpoints(ri resourceInfo, projectRoot string, baseIdx int) ([]scanner.Endpoint, []dtoRequest) {
	var endpoints []scanner.Endpoint
	var dtoReqs []dtoRequest
	for _, ep := range ri.endpoints {
		endpoint := buildEndpoint(ri, ep)
		epIdx := baseIdx + len(endpoints)
		reqs := collectDTORequests(ep, ri, projectRoot, epIdx)
		dtoReqs = append(dtoReqs, reqs...)
		endpoints = append(endpoints, endpoint)
	}
	return endpoints, dtoReqs
}
