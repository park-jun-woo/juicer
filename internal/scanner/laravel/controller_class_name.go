//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 배열 요소(Controller::class)에서 컨트롤러 클래스 이름을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func controllerClassName(elem *sitter.Node, src []byte) string {
	classAccess := findClassConstantAccess(elem)
	if classAccess == nil {
		return ""
	}
	nameNode := findChildByType(classAccess, "name")
	if nameNode == nil {
		return ""
	}
	return nodeText(nameNode, src)
}
