//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what 직접 자식 중 binary_expression이 있는지 확인한다
package supafunc

import sitter "github.com/smacker/go-tree-sitter"

func hasChildBinaryExpr(node *sitter.Node) bool {
	for i := 0; i < int(node.ChildCount()); i++ {
		if node.Child(i).Type() == "binary_expression" {
			return true
		}
	}
	return false
}
