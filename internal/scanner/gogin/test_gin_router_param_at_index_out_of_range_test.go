//ff:func feature=scan type=test control=sequence
//ff:what TestGinRouterParamAtIndex_OutOfRange 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestGinRouterParamAtIndex_OutOfRange(t *testing.T) {
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
	if got := ginRouterParamAtIndex(fn, nil, 5); got != "" {
		t.Fatalf("expected empty for out of range, got %q", got)
	}
}
