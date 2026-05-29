//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 노드가 .json field_expression 인지 판별한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func isJSONFieldExpr(node *sitter.Node, src []byte) bool {
	if node == nil || node.Type() != "field_expression" {
		return false
	}
	fieldID := findChildByType(node, "field_identifier")
	if fieldID == nil {
		return false
	}
	return nodeText(fieldID, src) == "json"
}
