//ff:func feature=scan type=test control=sequence
//ff:what fiberCtxParamName — 이름 없는 *fiber.Ctx 파라미터 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberCtxParamName_Unnamed(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					// no Names -> returns "_"
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
	if got := fiberCtxParamName(ft); got != "_" {
		t.Fatalf("expected '_' for unnamed ctx, got %q", got)
	}
}
