//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what TypeBox call_expression의 Type.X 메서드명(X)을 반환한다 (아니면 빈 문자열)
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func typeBoxCallName(call *sitter.Node, src []byte) string {
	if call == nil || call.Type() != "call_expression" {
		return ""
	}
	fn := findChildByType(call, "member_expression")
	if fn == nil {
		return ""
	}
	obj := findChildByType(fn, "identifier")
	if obj == nil || nodeText(obj, src) != "Type" {
		return ""
	}
	prop := fn.ChildByFieldName("property")
	if prop == nil {
		return ""
	}
	return nodeText(prop, src)
}
