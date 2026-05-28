//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what call_expression이 TypeBox Type.Object() 호출인지 확인한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func isTypeBoxObjectCall(call *sitter.Node, src []byte) bool {
	fn := findChildByType(call, "member_expression")
	if fn == nil {
		return false
	}
	obj := findChildByType(fn, "identifier")
	if obj == nil || nodeText(obj, src) != "Type" {
		return false
	}
	prop := fn.ChildByFieldName("property")
	if prop == nil {
		return false
	}
	return nodeText(prop, src) == "Object"
}
