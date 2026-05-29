//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 함수 본문에서 HttpResponse 상태 코드와 응답 종류를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractResponses(funcNode *sitter.Node, src []byte, sIdx structIndex, cache map[string][]scanner.Field) []scanner.Response {
	block := findChildByType(funcNode, "block")
	if block == nil {
		return nil
	}

	ctx := &responseCtx{block: block, src: src, sIdx: sIdx, cache: cache, seen: map[string]bool{}}
	walkNodes(block, func(n *sitter.Node) {
		captureResponse(n, ctx)
	})
	return ctx.responses
}
