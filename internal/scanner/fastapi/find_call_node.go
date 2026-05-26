//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 직접 자식 노드 중 call 타입의 첫 노드를 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findCallNode finds the first "call" node among direct children of the given node.
func findCallNode(node *sitter.Node) *sitter.Node {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "call" {
			return child
		}
	}
	return nil
}
