//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamNameInfo_NilInfo 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestGinCtxParamNameInfo_NilInfo(t *testing.T) {
	// info가 nil이면 AST 기반 fallback
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
	got := ginCtxParamNameInfo(ft, nil)
	if got != "c" {
		t.Fatalf("expected c (AST fallback), got %s", got)
	}
}
