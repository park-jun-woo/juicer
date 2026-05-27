//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 수집된 모델 요청 목록을 해석하여 엔드포인트에 필드를 채운다
package fastapi

import "github.com/park-jun-woo/juicer/internal/scanner"

// resolveAllModels resolves Pydantic model types and fills fields into endpoints.
func resolveAllModels(reqs []modelRequest, endpoints []scanner.Endpoint, files []fileInfo) {
	cache := make(map[string][]scanner.Field)
	globalModels := buildGlobalModelMap(files)
	enrichModelsWithInheritance(files, globalModels)

	for _, req := range reqs {
		fields := resolveModelFields(req, cache, globalModels)
		if fields == nil || req.epIdx >= len(endpoints) {
			continue
		}
		applyModelFields(&endpoints[req.epIdx], req.isBody, fields)
	}
}
