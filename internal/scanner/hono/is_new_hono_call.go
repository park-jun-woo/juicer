//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what new_expression 노드가 new Hono() 또는 new OpenAPIHono() 호출인지 확인한다
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
	name := nodeText(ident, src)
	return name == "Hono" || name == "OpenAPIHono"
}
