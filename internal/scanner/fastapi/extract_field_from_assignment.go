//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what assignment 노드에서 Pydantic 필드를 추출한다
package fastapi

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractFieldFromAssignment extracts a field from an assignment node.
// Handles: name: type = default_value
func extractFieldFromAssignment(assign *sitter.Node, src []byte) *pydanticField {
	ident := findChildByType(assign, "identifier")
	if ident == nil {
		return nil
	}
	name := nodeText(ident, src)
	if strings.HasPrefix(name, "__") {
		return nil
	}

	typeNode := findChildByType(assign, "type")
	typeName := ""
	if typeNode != nil {
		typeName = nodeText(typeNode, src)
	}

	return &pydanticField{
		name:       name,
		typeName:   unwrapNullable(typeName),
		hasDefault: true,
		nullable:   isNullableType(typeName),
	}
}
