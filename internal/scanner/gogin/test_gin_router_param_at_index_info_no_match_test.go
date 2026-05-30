//ff:func feature=scan type=test control=sequence
//ff:what TestGinRouterParamAtIndex_InfoNoMatch 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestGinRouterParamAtIndex_InfoNoMatch(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{Params: &ast.FieldList{List: []*ast.Field{
		{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "int"}},
	}}}}

	if got := ginRouterParamAtIndex(fn, &types.Info{}, 0); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
