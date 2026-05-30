//ff:func feature=scan type=test control=sequence topic=django
//ff:what firstExprStatement 테스트 헬퍼
package django

import sitter "github.com/smacker/go-tree-sitter"

func firstExprStatement(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "expression_statement" {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}
