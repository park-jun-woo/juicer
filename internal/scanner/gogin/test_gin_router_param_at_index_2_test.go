//ff:func feature=scan type=test control=sequence
//ff:what TestGinRouterParamAtIndex 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestGinRouterParamAtIndex(t *testing.T) {

	fn := &ast.FuncDecl{
		Name: &ast.Ident{Name: "f"},
		Type: &ast.FuncType{},
	}
	got := ginRouterParamAtIndex(fn, nil, 0)
	if got != "" {
		t.Fatal("expected empty for nil params")
	}

	fn2 := &ast.FuncDecl{
		Name: &ast.Ident{Name: "f"},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "int"}},
				},
			},
		},
	}
	got2 := ginRouterParamAtIndex(fn2, nil, 0)
	if got2 != "" {
		t.Fatal("expected empty for nil info")
	}
}
