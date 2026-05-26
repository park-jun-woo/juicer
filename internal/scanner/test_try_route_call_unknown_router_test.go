//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_UnknownRouter 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_UnknownRouter(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "unknown"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/test"`}, &ast.Ident{Name: "h"}},
	}
	routers := make(map[string]*routerInfo)
	fset := token.NewFileSet()
	_, _, ok := tryRouteCall(call, routers, "test.go", fset)
	if ok {
		t.Error("expected not ok for unknown router")
	}
}
