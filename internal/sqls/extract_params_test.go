package sqls

import (
	"go/ast"
	"testing"
)

func TestExtractParams_Nil(t *testing.T) {
	if extractParams(nil) != nil {
		t.Fatal("expected nil")
	}
}

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

func TestExtractParams_SkipContext(t *testing.T) {
	fields := &ast.FieldList{
		List: []*ast.Field{
			{Names: []*ast.Ident{{Name: "ctx"}}, Type: &ast.SelectorExpr{X: &ast.Ident{Name: "context"}, Sel: &ast.Ident{Name: "Context"}}},
		},
	}
	result := extractParams(fields)
	if len(result) != 0 {
		t.Fatalf("expected 0 (context skipped), got %d", len(result))
	}
}
