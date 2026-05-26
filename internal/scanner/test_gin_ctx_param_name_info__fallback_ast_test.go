//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamNameInfo_FallbackToAST 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestGinCtxParamNameInfo_FallbackToAST(t *testing.T) {
	// types.Info는 있지만 TypeOf가 nil을 반환하는 경우 → AST fallback
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "ctx"}},
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
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	got := ginCtxParamNameInfo(ft, info)
	if got != "ctx" {
		t.Fatalf("expected ctx (AST fallback), got %s", got)
	}
}
