//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what Rust struct의 필드를 파싱한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractStructFields(structNode *sitter.Node, src []byte) []scanner.Field {
	fieldList := findChildByType(structNode, "field_declaration_list")
	if fieldList == nil {
		return nil
	}

	var fields []scanner.Field
	var pendingAttrs []serdeAttr

	for i := 0; i < int(fieldList.ChildCount()); i++ {
		child := fieldList.Child(i)

		if child.Type() == "attribute_item" {
			attr := parseSerdeAttribute(child, src)
			if attr != nil {
				pendingAttrs = append(pendingAttrs, *attr)
			}
			continue
		}

		if child.Type() == "field_declaration" {
			f := parseOneField(child, src, pendingAttrs)
			if f != nil {
				fields = append(fields, *f)
			}
			pendingAttrs = nil
			continue
		}

		if child.Type() != "," {
			pendingAttrs = nil
		}
	}

	return fields
}

func parseOneField(fieldDecl *sitter.Node, src []byte, attrs []serdeAttr) *scanner.Field {
	// check serde(skip)
	for _, a := range attrs {
		if a.skip {
			return nil
		}
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

	jsonName := fieldName
	nullable := isOptionType(typeText)
	var actualType string
	if nullable {
		actualType = extractGenericInner(typeText)
	} else {
		actualType = typeText
	}

	// Apply serde attributes
	for _, a := range attrs {
		if a.rename != "" {
			jsonName = a.rename
		}
		if a.hasDefault {
			nullable = true
		}
	}

	oaType := rustTypeToOpenAPI(actualType)

	f := &scanner.Field{
		Name:     fieldName,
		Type:     oaType.Type,
		JSON:     jsonName,
		Nullable: nullable,
	}

	if oaType.Type == "array" && oaType.Items != "" {
		itemType := rustTypeToOpenAPI(oaType.Items)
		f.Fields = []scanner.Field{{
			Name: "items",
			Type: itemType.Type,
		}}
	}

	return f
}

func findFieldType(fieldDecl *sitter.Node) *sitter.Node {
	for i := 0; i < int(fieldDecl.ChildCount()); i++ {
		child := fieldDecl.Child(i)
		switch child.Type() {
		case "type_identifier", "generic_type", "primitive_type",
			"scoped_type_identifier", "reference_type":
			return child
		}
	}
	return nil
}
