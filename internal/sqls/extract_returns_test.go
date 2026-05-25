package sqls

import (
	"go/ast"
	"testing"
)

func TestExtractReturns_Nil(t *testing.T) {
	if extractReturns(nil) != nil {
		t.Fatal("expected nil")
	}
}

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
