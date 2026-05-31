//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 수집된 DTO 요청 목록을 해석하여 엔드포인트에 필드를 채운다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// resolveAllDTOs resolves all DTO types and fills fields into endpoints. It
// also recursively registers nested DTO/enum named types as separate component
// schemas, returning that schema map for ScanResult.Schemas.
func resolveAllDTOs(dtoReqs []dtoRequest, endpoints []scanner.Endpoint) map[string]any {
	cache := make(map[string][]scanner.Field)
	for _, dr := range dtoReqs {
		fields, err := resolveDTOFields(dr, cache)
		if err != nil || fields == nil || dr.epIdx >= len(endpoints) {
			continue
		}
		ep := &endpoints[dr.epIdx]
		if dr.isQuery {
			ensureRequest(ep)
			ep.Request.Query = fieldsToQueryParams(fields)
			continue
		}
		if dr.isBody && ep.Request != nil && ep.Request.Body != nil {
			ep.Request.Body.Fields = fields
		}
		if !dr.isBody && len(ep.Responses) > 0 {
			ep.Responses[0].Fields = fields
		}
	}
	schemas := make(map[string]any)
	registerNestedSchemas(dtoReqs, cache, schemas)
	return schemas
}
