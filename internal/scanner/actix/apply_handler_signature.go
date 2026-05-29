//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 핸들러 함수 시그니처에서 extractor와 응답을 추출해 엔드포인트에 반영한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyHandlerSignature(ep *scanner.Endpoint, funcNode *sitter.Node, src []byte, sIdx structIndex, cache map[string][]scanner.Field) {
	exts := extractExtractors(funcNode, src)
	applyExtractors(ep, exts, sIdx, cache)

	responses := extractResponses(funcNode, src)
	if len(responses) > 0 {
		ep.Responses = responses
	}
}
