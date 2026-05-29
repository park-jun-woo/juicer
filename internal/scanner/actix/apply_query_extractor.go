//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what web::Query extractor의 struct 필드를 쿼리 파라미터로 적용한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyQueryExtractor(ep *scanner.Endpoint, ext extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	fields := resolveStructFields(ext.typeName, sIdx, cache)
	if len(fields) == 0 {
		return
	}
	ensureRequest(ep)
	for _, f := range fields {
		ep.Request.Query = append(ep.Request.Query, scanner.Param{
			Name: f.JSON,
			Type: f.Type,
		})
	}
}
