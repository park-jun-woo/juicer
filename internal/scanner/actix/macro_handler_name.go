//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what function_item 노드에서 핸들러 함수 이름을 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func macroHandlerName(funcItem *sitter.Node, src []byte) string {
	nameNode := findChildByType(funcItem, "identifier")
	if nameNode == nil {
		return ""
	}
	return nodeText(nameNode, src)
}
