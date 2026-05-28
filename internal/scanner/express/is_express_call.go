//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression이 express() 호출인지 확인한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func isExpressCall(node *sitter.Node, src []byte) bool {
	if node.Type() != "call_expression" {
		return false
	}
	fn := findChildByType(node, "identifier")
	if fn != nil && nodeText(fn, src) == "express" {
		return true
	}
	return false
}
