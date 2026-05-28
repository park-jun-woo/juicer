//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 모든 컨트롤러에서 엔드포인트를 생성한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildAllEndpoints(controllers []controllerInfo, projectRoot string) ([]scanner.Endpoint, []dtoRequest) {
	var endpoints []scanner.Endpoint
	var dtoReqs []dtoRequest
	for _, ci := range controllers {
		eps, reqs := buildControllerEndpoints(ci, projectRoot, len(endpoints))
		endpoints = append(endpoints, eps...)
		dtoReqs = append(dtoReqs, reqs...)
	}
	return endpoints, dtoReqs
}
