//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 모든 리소스에서 엔드포인트를 생성한다
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildAllEndpoints(resources []resourceInfo, projectRoot string) ([]scanner.Endpoint, []dtoRequest) {
	var endpoints []scanner.Endpoint
	var dtoReqs []dtoRequest
	for _, ri := range resources {
		eps, reqs := buildResourceEndpoints(ri, projectRoot, len(endpoints))
		endpoints = append(endpoints, eps...)
		dtoReqs = append(dtoReqs, reqs...)
	}
	return endpoints, dtoReqs
}
