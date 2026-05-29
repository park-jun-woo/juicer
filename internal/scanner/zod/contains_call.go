//ff:func feature=scan type=extract control=sequence topic=zod
//ff:what AST 노드에 z.method() 호출이 포함되어 있는지 확인한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

// ContainsCall — z.xxx() 호출 포함 여부
func ContainsCall(node *sitter.Node, src []byte) bool {
	found := false
	walkNodes(node, func(n *sitter.Node) {
		if found {
			return
		}
		if n.Type() == "member_expression" {
			obj := findChildByType(n, "identifier")
			if obj != nil && nodeText(obj, src) == "z" {
				found = true
			}
		}
	})
	return found
}
