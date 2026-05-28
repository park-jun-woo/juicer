//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 모든 DTO 타입을 해석하여 엔드포인트에 필드를 할당한다
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

func resolveAllDTOs(dtoReqs []dtoRequest, endpoints []scanner.Endpoint, projectRoot string) {
	cache := make(map[string][]scanner.Field)
	for _, dr := range dtoReqs {
		if dr.epIdx >= len(endpoints) {
			continue
		}
		fields := resolveDTOType(dr, projectRoot, cache)
		if fields == nil {
			continue
		}
		assignDTOFields(dr, &endpoints[dr.epIdx], fields)
	}
}
