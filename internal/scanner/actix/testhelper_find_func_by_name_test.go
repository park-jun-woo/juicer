//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findFuncByName 테스트 헬퍼
package actix

import sitter "github.com/smacker/go-tree-sitter"

// findFuncByName returns the first function_item node whose name matches.
func findFuncByName(root *sitter.Node, src []byte, name string) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "function_item" || n.Type() == "function_signature_item" {
			id := findChildByType(n, "identifier")
			if id != nil && nodeText(id, src) == name {
				found = n
				return
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}
