//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what z.object({...}) AST를 scanner.Field 슬라이스로 변환한다
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// ParseSchema — z.object() → Field 슬라이스
func ParseSchema(node *sitter.Node, src []byte) []scanner.Field {
	if node == nil {
		return nil
	}
	calls := FindObjectCalls(node, src)
	if len(calls) > 0 {
		return ParseObjectArgs(calls[0], src)
	}
	if IsObjectCall(node, src) {
		return ParseObjectArgs(node, src)
	}
	return nil
}
