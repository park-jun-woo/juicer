//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what type_annotation 노드에서 타입 문자열을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractTypeAnnotation returns the type string from a type_annotation node.
func extractTypeAnnotation(node *sitter.Node, src []byte) string {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() != ":" {
			return nodeText(child, src)
		}
	}
	return ""
}
