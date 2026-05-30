//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_NotMethod 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_NotMethod(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("main.go", 1, 100)

	routers := map[string]*routerInfo{
		"app": {},
	}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "app"},
			Sel: &ast.Ident{Name: "Use"},
		},
		Args: []ast.Expr{
			&ast.Ident{Name: "logger"},
		},
	}

	_, _, ok := tryRouteCall(call, routers, "main.go", fset)
	if ok {
		t.Fatal("expected false — Use is not a route method")
	}
}
