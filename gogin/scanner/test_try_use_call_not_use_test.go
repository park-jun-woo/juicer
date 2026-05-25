//ff:func feature=scan type=extract control=sequence
//ff:what TestTryUseCall_NotUse 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_NotUse(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
	}
	routers := make(map[string]*routerInfo)
	tryUseCall(call, routers)
}
