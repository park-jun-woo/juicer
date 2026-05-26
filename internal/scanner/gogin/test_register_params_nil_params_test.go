//ff:func feature=scan type=extract control=sequence
//ff:what TestRegisterParams_NilParams 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestRegisterParams_NilParams(t *testing.T) {
	fn := &ast.FuncDecl{
		Name: &ast.Ident{Name: "test"},
		Type: &ast.FuncType{Params: nil},
	}
	routers := make(map[string]*routerInfo)
	registerParams(fn, "gin", routers)
	if len(routers) != 0 {
		t.Errorf("expected 0 routers, got %d", len(routers))
	}
}
