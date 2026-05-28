//ff:func feature=scan type=test control=sequence
//ff:what TestEchoCtxParamName_Found 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestEchoCtxParamName_Found(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "c"}},
					Type: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "echo"},
						Sel: &ast.Ident{Name: "Context"},
					},
				},
			},
		},
	}
	got := echoCtxParamName(ft)
	if got != "c" {
		t.Fatalf("expected c, got %s", got)
	}
}
