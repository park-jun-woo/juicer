//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what web::Json extractor를 요청 본문(body)으로 적용한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyJSONExtractor(ep *scanner.Endpoint, ext extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	ensureRequest(ep)
	ep.Request.Body = &scanner.Body{
		VarName:  "body",
		Method:   "json",
		TypeName: ext.typeName,
	}
	fields := resolveStructFields(ext.typeName, sIdx, cache)
	if len(fields) > 0 {
		ep.Request.Body.Fields = fields
	}
}
