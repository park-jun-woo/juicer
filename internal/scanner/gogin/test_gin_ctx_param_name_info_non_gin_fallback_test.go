//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamNameInfo_NonGinFallback 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestGinCtxParamNameInfo_NonGinFallback(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "int"}},
			},
		},
	}
	if got := ginCtxParamNameInfo(ft, &types.Info{}); got != "" {
		t.Fatalf("non-gin: got %q", got)
	}
}
