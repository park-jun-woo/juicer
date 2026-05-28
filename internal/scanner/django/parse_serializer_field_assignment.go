//ff:func feature=scan type=extract control=sequence topic=django
//ff:what Serializer 필드 대입문에서 필드 정보를 추출한다
package django

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// parseSerializerFieldAssignment parses a field assignment in a serializer body.
func parseSerializerFieldAssignment(stmtNode *sitter.Node, src []byte) *scanner.Field {
	assign := findChildByType(stmtNode, "assignment")
	if assign == nil {
		return nil
	}
	nameNode := findChildByType(assign, "identifier")
	if nameNode == nil {
		return nil
	}
	fieldName := nodeText(nameNode, src)
	if fieldName == "Meta" {
		return nil
	}
	callNode := findChildByType(assign, "call")
	if callNode == nil {
		return nil
	}
	fieldType := extractDRFFieldType(callNode, src)
	if fieldType == "" {
		return nil
	}
	oaType := drfFieldToOpenAPI(fieldType)
	field := &scanner.Field{
		Name: fieldName,
		Type: oaType.Type,
		JSON: fieldName,
	}
	args := findChildByType(callNode, "argument_list")
	if args != nil {
		applySerializerFieldConstraints(field, args, src, fieldType)
	}
	return field
}
