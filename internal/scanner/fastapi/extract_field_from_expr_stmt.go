//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what expression_statement 에서 Pydantic 필드를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractFieldFromExprStmt extracts a field from an expression statement.
// Handles patterns like: name: str  or  name: str = "default"
func extractFieldFromExprStmt(stmt *sitter.Node, src []byte) *pydanticField {
	assign := findChildByType(stmt, "assignment")
	if assign != nil {
		return extractFieldFromAssignment(assign, src)
	}

	ident := findChildByType(stmt, "identifier")
	ann := findChildByType(stmt, "type")
	if ident != nil && ann != nil {
		name := nodeText(ident, src)
		typeName := nodeText(ann, src)
		return &pydanticField{
			name:     name,
			typeName: unwrapNullable(typeName),
			nullable: isNullableType(typeName),
		}
	}
	return nil
}
