//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what struct_expression 노드에서 type_identifier 텍스트(타입명)를 반환한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func structExprTypeName(structExpr *sitter.Node, src []byte) string {
	typeID := findChildByType(structExpr, "type_identifier")
	if typeID == nil {
		return ""
	}
	return nodeText(typeID, src)
}
