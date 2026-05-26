//ff:func feature=scan type=test control=sequence
//ff:what TestGinRouterParamAtIndex_UnnamedParam 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestGinRouterParamAtIndex_UnnamedParam(t *testing.T) {
	fn := &ast.FuncDecl{
		Name: &ast.Ident{Name: "f"},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: &ast.Ident{Name: "int"},
					},
				},
			},
		},
	}
	if got := ginRouterParamAtIndex(fn, nil, 0); got != "" {
		t.Fatalf("expected empty for unnamed param with nil info, got %q", got)
	}
}
