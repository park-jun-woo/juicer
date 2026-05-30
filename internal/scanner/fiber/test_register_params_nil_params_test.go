//ff:func feature=scan type=test control=sequence
//ff:what TestRegisterParams_NilParams 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestRegisterParams_NilParams(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	routers := map[string]*routerInfo{}
	registerParams(fn, "fiber", routers)
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}
