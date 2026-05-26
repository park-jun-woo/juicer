//ff:func feature=scan type=test control=sequence
//ff:what TestGinRouterParamAtIndex_NilParams 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestGinRouterParamAtIndex_NilParams(t *testing.T) {
	fn := &ast.FuncDecl{
		Name: &ast.Ident{Name: "f"},
		Type: &ast.FuncType{},
	}
	if got := ginRouterParamAtIndex(fn, nil, 0); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
