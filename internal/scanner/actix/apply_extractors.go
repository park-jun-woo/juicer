//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what extractor 목록을 종류별로 엔드포인트 요청에 적용한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyExtractors(ep *scanner.Endpoint, exts []extractorInfo, sIdx structIndex, cache map[string][]scanner.Field) {
	for _, ext := range exts {
		applyExtractor(ep, ext, sIdx, cache)
	}
}
