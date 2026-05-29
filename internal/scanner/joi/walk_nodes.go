//ff:func feature=scan type=extract control=iteration dimension=1 topic=joi
//ff:what tree-sitter 노드를 재귀적으로 순회하며 콜백을 호출한다
package joi

import sitter "github.com/smacker/go-tree-sitter"

func walkNodes(node *sitter.Node, fn func(*sitter.Node)) {
	fn(node)
	for i := 0; i < int(node.ChildCount()); i++ {
		walkNodes(node.Child(i), fn)
	}
}
