//ff:func feature=scan type=extract control=iteration dimension=1 topic=joi
//ff:what 직계 자식 중 지정 타입의 모든 노드를 수집한다
package joi

import sitter "github.com/smacker/go-tree-sitter"

func childrenOfType(node *sitter.Node, typeName string) []*sitter.Node {
	var result []*sitter.Node
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == typeName {
			result = append(result, child)
		}
	}
	return result
}
