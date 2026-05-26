//ff:func feature=sql type=test control=sequence
//ff:what TestExtractReturns_WithTypesCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestExtractReturns_WithTypesCov(t *testing.T) {
	fields := &ast.FieldList{List: []*ast.Field{
		{Type: &ast.Ident{Name: "error"}},
		{Names: []*ast.Ident{{Name: "result"}}, Type: &ast.Ident{Name: "string"}},
	}}
	returns := extractReturns(fields)
	if len(returns) != 2 {
		t.Fatalf("expected 2, got %d", len(returns))
	}
}
