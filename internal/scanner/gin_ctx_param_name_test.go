package scanner

import (
	"go/ast"
	"testing"
)

func TestGinCtxParamName_WithGinCtx(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "c"}},
					Type: &ast.StarExpr{
						X: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "gin"},
							Sel: &ast.Ident{Name: "Context"},
						},
					},
				},
			},
		},
	}
	got := ginCtxParamName(ft)
	if got != "c" {
		t.Fatalf("expected c, got %s", got)
	}
}

func TestGinCtxParamName_NoGinCtx(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "int"}},
			},
		},
	}
	got := ginCtxParamName(ft)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}

func TestGinCtxParamName_NilParams(t *testing.T) {
	ft := &ast.FuncType{}
	got := ginCtxParamName(ft)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
