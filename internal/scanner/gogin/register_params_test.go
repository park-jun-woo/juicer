//ff:func feature=scan type=test control=sequence
//ff:what TestRegisterParams_EmptyParams 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestRegisterParams_EmptyParams(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	routers := map[string]*routerInfo{}
	registerParams(fn, "gin", routers)
	if len(routers) != 0 {
		t.Fatal("expected no routers")
	}
}

