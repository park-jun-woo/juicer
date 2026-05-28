//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression이 res.status(N) 호출인지 판별한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func isResStatusCall(call *sitter.Node, src []byte) bool {
	if call.Type() != "call_expression" {
		return false
	}
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return false
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return false
	}
	if nodeText(prop, src) != "status" {
		return false
	}
	obj := mem.ChildByFieldName("object")
	if obj == nil {
		return false
	}
	return obj.Type() == "identifier" && resObjectNames[nodeText(obj, src)]
}
