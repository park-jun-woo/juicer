//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_PathParams 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_PathParams(t *testing.T) {
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
			&ast.BasicLit{Kind: token.STRING, Value: `"/users/:id"`},
			&ast.Ident{Name: "getUser"},
		},
	}

	ep, _, ok := tryRouteCall(call, routers, "main.go", fset)
	if !ok {
		t.Fatal("expected ok")
	}
	if ep.Request == nil {
		t.Fatal("expected non-nil Request")
	}
	if len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "id" {
		t.Fatalf("expected path param 'id', got %v", ep.Request.PathParams)
	}
}
