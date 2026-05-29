//ff:func feature=scan type=extract control=iteration dimension=1 topic=zod
//ff:what 직계 자식 중 지정 타입의 첫 노드를 반환한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

func findChildByType(node *sitter.Node, typeName string) *sitter.Node {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == typeName {
			return child
		}
	}
	return nil
}
