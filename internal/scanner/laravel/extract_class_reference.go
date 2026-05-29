//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Controller::class 표현식에서 클래스 이름을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractClassReference extracts a class name from Controller::class expression.
func extractClassReference(node *sitter.Node, src []byte) string {
	classAccess := findClassConstantAccess(node)
	if classAccess == nil {
		return ""
	}
	nameNode := findChildByType(classAccess, "name")
	if nameNode == nil {
		return ""
	}
	return nodeText(nameNode, src)
}
