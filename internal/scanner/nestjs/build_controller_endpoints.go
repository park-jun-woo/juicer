//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 단일 컨트롤러에서 엔드포인트와 DTO 요청 목록을 생성한다
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// buildControllerEndpoints builds endpoints for a single controller.
// baseIdx is the current length of the global endpoints slice for correct DTO indexing.
func buildControllerEndpoints(globalPrefix string, cwf controllerWithFile, baseIdx int) ([]scanner.Endpoint, []dtoRequest) {
	ci := cwf.info
	var endpoints []scanner.Endpoint
	var dtoReqs []dtoRequest
	for _, ep := range ci.endpoints {
		endpoint := buildEndpoint(globalPrefix, ci, ep)
		epIdx := baseIdx + len(endpoints)
		reqs := collectDTORequests(ep, ci.imports, cwf.absFile, epIdx)
		dtoReqs = append(dtoReqs, reqs...)
		endpoints = append(endpoints, endpoint)
	}
	return endpoints, dtoReqs
}
