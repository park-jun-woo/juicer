//ff:func feature=scan type=test control=sequence
//ff:what TestGinRouterParamAtIndex_NilInfo 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestGinRouterParamAtIndex_NilInfo(t *testing.T) {
	fn := &ast.FuncDecl{
		Name: &ast.Ident{Name: "f"},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: "rg"}},
						Type:  &ast.Ident{Name: "RouterGroup"},
					},
				},
			},
		},
	}
	if got := ginRouterParamAtIndex(fn, nil, 0); got != "" {
		t.Fatalf("expected empty with nil info, got %q", got)
	}
}
