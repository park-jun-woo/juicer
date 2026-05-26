//ff:func feature=scan type=test control=sequence
//ff:what TestGinRouterParamAtIndex_NonRouterType 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestGinRouterParamAtIndex_NonRouterType(t *testing.T) {
	fn := &ast.FuncDecl{
		Name: &ast.Ident{Name: "f"},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: "x"}},
						Type:  &ast.Ident{Name: "int"},
					},
				},
			},
		},
	}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
	}
	if got := ginRouterParamAtIndex(fn, info, 0); got != "" {
		t.Fatalf("expected empty for non-router type, got %q", got)
	}
}
