//ff:func feature=scan type=extract control=sequence
//ff:what TestGinCtxParamName_NoGinCtx 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestGinCtxParamName_NoGinCtx(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "int"}},
			},
		},
	}
	got := ginCtxParamName(ft)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
