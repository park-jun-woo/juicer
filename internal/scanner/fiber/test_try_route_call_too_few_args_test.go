//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_TooFewArgs 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_TooFewArgs(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("main.go", 1, 100)

	routers := map[string]*routerInfo{
		"app": {},
	}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "app"},
			Sel: &ast.Ident{Name: "Get"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"/users"`},
		},
	}

	_, _, ok := tryRouteCall(call, routers, "main.go", fset)
	if ok {
		t.Fatal("expected false — only 1 arg, need at least 2")
	}
}
