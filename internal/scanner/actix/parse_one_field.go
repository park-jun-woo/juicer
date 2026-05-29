//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what field_declaration 하나를 scanner.Field로 변환한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func parseOneField(fieldDecl *sitter.Node, src []byte, attrs []serdeAttr) *scanner.Field {
	if serdeAttrsSkip(attrs) {
		return nil
	}

	nameNode := findChildByType(fieldDecl, "field_identifier")
	if nameNode == nil {
		return nil
	}
	fieldName := nodeText(nameNode, src)

	typeNode := findFieldType(fieldDecl)
	if typeNode == nil {
		return nil
	}
	typeText := nodeText(typeNode, src)

	nullable := isOptionType(typeText)
	actualType := typeText
	if nullable {
		actualType = extractGenericInner(typeText)
	}

	jsonName, nullable := applySerdeAttrs(attrs, fieldName, nullable)
	oaType := rustTypeToOpenAPI(actualType)

	f := &scanner.Field{
		Name:     fieldName,
		Type:     oaType.Type,
		JSON:     jsonName,
		Nullable: nullable,
	}
	attachArrayItems(f, oaType)
	return f
}
