//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamNameInfo_InfoFallback 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestGinCtxParamNameInfo_InfoFallback(t *testing.T) {

	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "c"}},
					Type: &ast.StarExpr{X: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "gin"},
						Sel: &ast.Ident{Name: "Context"},
					}},
				},
			},
		},
	}
	info := &types.Info{}
	if got := ginCtxParamNameInfo(ft, info); got != "c" {
		t.Fatalf("fallback: got %q, want c", got)
	}
}
