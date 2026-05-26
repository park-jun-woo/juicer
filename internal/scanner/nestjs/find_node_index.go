//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 부모 노드에서 자식 노드의 인덱스를 찾는다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// findNodeIndex returns the index of node within parent's children.
func findNodeIndex(parent, node *sitter.Node) int {
	for i := 0; i < int(parent.ChildCount()); i++ {
		if parent.Child(i) == node {
			return i
		}
	}
	return -1
}
