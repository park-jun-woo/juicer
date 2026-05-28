//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression이 router.route() 호출인지 확인한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func isRouteCall(node *sitter.Node, src []byte, routers map[string]bool) bool {
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
	if !routers[nodeText(obj, src)] {
		return false
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return false
	}
	return nodeText(prop, src) == "route"
}
