//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_UnknownRouter 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_UnknownRouter(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("main.go", 1, 100)

	routers := map[string]*routerInfo{}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "unknown"},
			Sel: &ast.Ident{Name: "Get"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"/users"`},
			&ast.Ident{Name: "handler"},
		},
	}

	_, _, ok := tryRouteCall(call, routers, "main.go", fset)
	if ok {
		t.Fatal("expected false — unknown router")
	}
}
