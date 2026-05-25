//ff:func feature=scan type=extract control=sequence
//ff:what TestTryUseCall_UnknownRouter 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_UnknownRouter(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "unknown"}, Sel: &ast.Ident{Name: "Use"}},
		Args: []ast.Expr{&ast.Ident{Name: "mw"}},
	}
	routers := make(map[string]*routerInfo)
	tryUseCall(call, routers)
}
