//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what web::Path extractor를 path 파라미터(원시 타입 또는 struct 필드)로 적용한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyPathExtractor(ep *scanner.Endpoint, ext extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	oaType := rustTypeToOpenAPI(ext.typeName)
	if oaType.Type != "object" {
		applyPrimitivePathType(ep, oaType.Type)
		return
	}
	fields := resolveStructFields(ext.typeName, sIdx, cache)
	if len(fields) > 0 {
		applyStructPathParams(ep, fields)
	}
}
