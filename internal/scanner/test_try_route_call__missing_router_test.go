//ff:func feature=scan type=extract control=sequence
//ff:what TestTryRouteCall_MissingRouter 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryRouteCall_MissingRouter(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/test"`}, &ast.Ident{Name: "h"}},
	}
	_, ok := tryRouteCall(call, map[string]*routerInfo{}, "", nil)
	if ok {
		t.Fatal("expected false for unknown router")
	}
}
