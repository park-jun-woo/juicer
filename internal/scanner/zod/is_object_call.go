//ff:func feature=scan type=extract control=sequence topic=zod
//ff:what call_expression이 z.object() 호출인지 확인한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

// IsObjectCall — z.object() 호출 여부 판별
func IsObjectCall(node *sitter.Node, src []byte) bool {
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
	if nodeText(obj, src) != "z" {
		return false
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return false
	}
	return nodeText(prop, src) == "object"
}
