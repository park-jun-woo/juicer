//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 클래스 선언의 이름이 주어진 이름과 일치하는지(이름 노드 없으면 통과) 보고한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func classNameMatches(cls *sitter.Node, src []byte, className string) bool {
	nameNode := findChildByType(cls, "name")
	if nameNode == nil {
		return true
	}
	return nodeText(nameNode, src) == className
}
