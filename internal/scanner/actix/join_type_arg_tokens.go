//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what type_arguments의 토큰들(꺾쇠 제외)을 이어붙여 튜플/복합 타입 텍스트를 만든다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func joinTypeArgTokens(typeArgs *sitter.Node, src []byte) string {
	result := ""
	for i := 0; i < int(typeArgs.ChildCount()); i++ {
		child := typeArgs.Child(i)
		t := child.Type()
		if t == "<" || t == ">" {
			continue
		}
		result += nodeText(child, src)
	}
	return result
}
