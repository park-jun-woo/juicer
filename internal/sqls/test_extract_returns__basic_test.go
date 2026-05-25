//ff:func feature=sql type=parse control=sequence
//ff:what TestExtractReturns_Basic 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestExtractReturns_Basic(t *testing.T) {
	fields := &ast.FieldList{
		List: []*ast.Field{
			{Type: &ast.Ident{Name: "error"}},
		},
	}
	result := extractReturns(fields)
	if len(result) != 1 || result[0] != "error" {
		t.Fatalf("expected 'error', got %v", result)
	}
}
