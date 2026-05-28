//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 어노테이션 노드에서 이름을 추출한다
package spring

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func annotationName(ann *sitter.Node, src []byte) string {
	for i := 0; i < int(ann.ChildCount()); i++ {
		child := ann.Child(i)
		if child.Type() == "identifier" || child.Type() == "scoped_identifier" {
			text := nodeText(child, src)
			parts := strings.Split(text, ".")
			return parts[len(parts)-1]
		}
	}
	return ""
}
