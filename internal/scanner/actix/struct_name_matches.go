//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what struct 노드의 타입명이 주어진 이름과 일치하는지 판별한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func structNameMatches(structNode *sitter.Node, src []byte, typeName string) bool {
	nameNode := findChildByType(structNode, "type_identifier")
	if nameNode == nil {
		return false
	}
	return nodeText(nameNode, src) == typeName
}
