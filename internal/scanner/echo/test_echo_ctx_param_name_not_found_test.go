//ff:func feature=scan type=test control=sequence
//ff:what TestEchoCtxParamName_NotFound 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestEchoCtxParamName_NotFound(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "name"}},
					Type:  &ast.Ident{Name: "string"},
				},
			},
		},
	}
	got := echoCtxParamName(ft)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
