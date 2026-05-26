//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamName_WithGinCtx 테스트
package gogin

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

	// nil params
	if ginCtxParamName(&ast.FuncType{}) != "" {
		t.Fatal("nil params")
	}

	// unnamed gin.Context param
	ft2 := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{{
				Type: &ast.StarExpr{X: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Context"}}},
			}},
		},
	}
	if ginCtxParamName(ft2) != "_" {
		t.Fatal("unnamed")
	}

	// no gin.Context param
	ft3 := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "int"}}},
		},
	}
	if ginCtxParamName(ft3) != "" {
		t.Fatal("no match")
	}
}

