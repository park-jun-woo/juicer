//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what Serializer 클래스 body에서 필드 정의를 추출한다
package django

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractSerializerFields extracts field definitions from a Serializer class body.
func extractSerializerFields(body *sitter.Node, src []byte) []scanner.Field {
	var fields []scanner.Field
	for _, stmtNode := range childrenOfType(body, "expression_statement") {
		f := parseSerializerFieldAssignment(stmtNode, src)
		if f != nil {
			fields = append(fields, *f)
		}
	}
	return fields
}
