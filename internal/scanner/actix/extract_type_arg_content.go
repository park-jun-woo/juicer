//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what type_arguments 노드에서 내부 타입 텍스트를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractTypeArgContent(typeArgs *sitter.Node, src []byte) string {
	var parts []string
	for i := 0; i < int(typeArgs.ChildCount()); i++ {
		child := typeArgs.Child(i)
		if child.IsNamed() {
			parts = append(parts, nodeText(child, src))
		}
	}
	if len(parts) == 1 {
		return parts[0]
	}
	return joinTypeArgTokens(typeArgs, src)
}
