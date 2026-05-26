//ff:func feature=scan type=extract control=sequence
//ff:what TestTryUseCall_MissingRouter 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_MissingRouter(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Use"}},
		Args: []ast.Expr{&ast.Ident{Name: "mw"}},
	}
	tryUseCall(call, map[string]*routerInfo{})
}
