//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamName_NoGinCtxCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestGinCtxParamName_NoGinCtxCov(t *testing.T) {
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
