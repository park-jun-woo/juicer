//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what 블록 자식 노드에서 필드 추출을 시도한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// tryExtractField tries to extract a field from a block child node.
func tryExtractField(child *sitter.Node, src []byte) *pydanticField {
	switch child.Type() {
	case "expression_statement":
		return extractFieldFromExprStmt(child, src)
	case "assignment":
		return extractFieldFromAssignment(child, src)
	}
	return nil
}
