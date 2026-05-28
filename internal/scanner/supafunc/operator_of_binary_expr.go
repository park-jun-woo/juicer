//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what binary_expression 노드의 연산자 텍스트를 반환한다 (left, right를 제외한 중간 자식)
package supafunc

import sitter "github.com/smacker/go-tree-sitter"

func operatorOfBinaryExpr(n *sitter.Node, src []byte) string {
	left := n.ChildByFieldName("left")
	right := n.ChildByFieldName("right")
	for i := 0; i < int(n.ChildCount()); i++ {
		child := n.Child(i)
		if child == left || child == right {
			continue
		}
		text := nodeText(child, src)
		if text != "" {
			return text
		}
	}
	return ""
}
