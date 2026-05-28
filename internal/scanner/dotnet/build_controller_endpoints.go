//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 단일 컨트롤러의 엔드포인트를 생성한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildControllerEndpoints(ci controllerInfo, projectRoot string, baseIdx int) ([]scanner.Endpoint, []dtoRequest) {
	var endpoints []scanner.Endpoint
	var dtoReqs []dtoRequest
	for _, ep := range ci.endpoints {
		endpoint := buildEndpoint(ci, ep)
		epIdx := baseIdx + len(endpoints)
		reqs := collectDTORequests(ep, ci, projectRoot, epIdx)
		dtoReqs = append(dtoReqs, reqs...)
		endpoints = append(endpoints, endpoint)
	}
	return endpoints, dtoReqs
}
