//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what new_expression 노드가 new Hono() 호출인지 확인한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func isNewHonoCall(node *sitter.Node, src []byte) bool {
	if node.Type() != "new_expression" {
		return false
	}
	ident := findChildByType(node, "identifier")
	if ident == nil {
		return false
	}
	return nodeText(ident, src) == "Hono"
}
