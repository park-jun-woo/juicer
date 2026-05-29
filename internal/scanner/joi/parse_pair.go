//ff:func feature=scan type=parse control=sequence topic=joi
//ff:what object pair 노드에서 키명과 Joi 체인을 파싱하여 Field를 생성한다
package joi

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// ParsePair — pair → Field
func ParsePair(pair *sitter.Node, src []byte) *scanner.Field {
	keyNode := pair.ChildByFieldName("key")
	if keyNode == nil {
		return nil
	}
	key := unquoteJoi(nodeText(keyNode, src))
	valueNode := pair.ChildByFieldName("value")
	if valueNode == nil {
		return nil
	}
	f := ParseChain(valueNode, src)
	f.Name = key
	f.JSON = key
	return &f
}
