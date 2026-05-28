//ff:func feature=scan type=test control=sequence
//ff:what TestFiberCtxParamName_Found 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberCtxParamName_Found(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "c"}},
					Type: &ast.StarExpr{
						X: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "fiber"},
							Sel: &ast.Ident{Name: "Ctx"},
						},
					},
				},
			},
		},
	}
	got := fiberCtxParamName(ft)
	if got != "c" {
		t.Fatalf("expected c, got %s", got)
	}
}
