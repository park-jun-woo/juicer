//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what AST 서브트리를 순회하며 콜백을 호출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// walkNodes calls fn for every node in the subtree (pre-order).
func walkNodes(node *sitter.Node, fn func(*sitter.Node)) {
	fn(node)
	for i := 0; i < int(node.ChildCount()); i++ {
		walkNodes(node.Child(i), fn)
	}
}
