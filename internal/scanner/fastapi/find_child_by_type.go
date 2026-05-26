//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 직접 자식 노드 중 지정 타입의 첫 노드를 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findChildByType returns the first direct child with the given type.
func findChildByType(node *sitter.Node, typeName string) *sitter.Node {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == typeName {
			return child
		}
	}
	return nil
}
