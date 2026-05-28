//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression이 express.Router() 호출인지 확인한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func isExpressRouterCall(node *sitter.Node, src []byte) bool {
	if node.Type() != "call_expression" {
		return false
	}
	mem := findChildByType(node, "member_expression")
	if mem == nil {
		return false
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return false
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return false
	}
	return nodeText(obj, src) == "express" && nodeText(prop, src) == "Router"
}
