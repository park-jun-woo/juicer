//ff:func feature=sql type=parse control=sequence
//ff:what TestExtractParams_Basic 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestExtractParams_Basic(t *testing.T) {
	fields := &ast.FieldList{
		List: []*ast.Field{
			{Names: []*ast.Ident{{Name: "id"}}, Type: &ast.Ident{Name: "int"}},
		},
	}
	result := extractParams(fields)
	if len(result) != 1 || result[0] != "id int" {
		t.Fatalf("expected 'id int', got %v", result)
	}
}
