//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what TypeBox Type.Object 인자 객체 노드를 scanner.Field 슬라이스로 변환한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func typeBoxObjectToFields(objNode *sitter.Node, src []byte) []scanner.Field {
	if objNode == nil || objNode.Type() != "object" {
		return nil
	}
	var fields []scanner.Field
	for _, pair := range childrenOfType(objNode, "pair") {
		f := typeBoxPairToField(pair, src)
		if f != nil {
			fields = append(fields, *f)
		}
	}
	return fields
}
