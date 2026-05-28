//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 직접 자식 노드 중 지정 타입만 수집한다
package express

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
