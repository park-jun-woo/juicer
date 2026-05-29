//ff:func feature=scan type=parse control=iteration dimension=1 topic=zod
//ff:what z.string().email().min(1) 등 Zod 메서드 체인을 파싱하여 Field를 생성한다
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// ParseChain — 메서드 체인 → Field
func ParseChain(node *sitter.Node, src []byte) scanner.Field {
	var f scanner.Field
	methods := CollectChainMethods(node, src)
	for _, m := range methods {
		ApplyMethod(&f, m)
	}
	return f
}
