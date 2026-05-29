//ff:func feature=scan type=extract control=selection topic=actix
//ff:what extractor 하나를 종류(path/json/query/form)에 따라 엔드포인트에 적용한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyExtractor(ep *scanner.Endpoint, ext extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	switch ext.kind {
	case "path":
		applyPathExtractor(ep, ext, sIdx, cache)
	case "json":
		applyJSONExtractor(ep, ext, sIdx, cache)
	case "query":
		applyQueryExtractor(ep, ext, sIdx, cache)
	case "form":
		applyFormExtractor(ep, ext, sIdx, cache)
	}
}
