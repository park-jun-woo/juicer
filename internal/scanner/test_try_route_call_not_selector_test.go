//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_NotSelector 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_NotSelector(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "fn"}}
	routers := make(map[string]*routerInfo)
	fset := token.NewFileSet()
	_, ok := tryRouteCall(call, routers, "test.go", fset)
	if ok {
		t.Error("expected not ok for non-selector")
	}
}
