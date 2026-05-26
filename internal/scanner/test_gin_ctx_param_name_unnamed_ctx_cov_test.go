//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamName_UnnamedCtxCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestGinCtxParamName_UnnamedCtxCov(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
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
	if got != "_" {
		t.Fatalf("expected _, got %s", got)
	}
}
