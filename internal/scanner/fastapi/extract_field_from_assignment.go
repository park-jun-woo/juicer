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

	f := &pydanticField{
		name:       name,
		typeName:   unwrapNullable(typeName),
		hasDefault: true,
		nullable:   isNullableType(typeName),
	}

	// Extract Field(...) constraints: ge, le, min_length, max_length
	// Field(...) 호출이 있으면 hasDefault를 false로 리셋 후 default 키워드 여부로 재결정
	if hasFieldCall(assign, src) {
		f.hasDefault = false
	}
	extractFieldCallConstraints(assign, src, f)

	return f
}
