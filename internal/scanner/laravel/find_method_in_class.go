//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 클래스 선언 안에서 지정 이름의 method_declaration을 찾는다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findMethodInClass(cls *sitter.Node, src []byte, methodName string) *sitter.Node {
	declList := findChildByType(cls, "declaration_list")
	if declList == nil {
		return nil
	}
	for _, method := range childrenOfType(declList, "method_declaration") {
		mName := findChildByType(method, "name")
		if mName != nil && nodeText(mName, src) == methodName {
			return method
		}
	}
	return nil
}
