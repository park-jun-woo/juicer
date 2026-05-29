//ff:func feature=scan type=parse control=iteration dimension=1 topic=joi
//ff:what Joi.string().required().email() 등 메서드 체인을 파싱하여 Field를 생성한다
package joi

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
	if f.Type == "" {
		f.Type = "string"
	}
	return f
}
