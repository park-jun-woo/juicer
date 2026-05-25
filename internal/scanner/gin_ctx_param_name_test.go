//ff:func feature=scan type=extract control=sequence
//ff:what TestGinCtxParamName_WithGinCtx 테스트
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
