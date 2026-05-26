//ff:func feature=scan type=test control=sequence
//ff:what TestTryUseCall_NoRouterCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_NoRouterCov(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "unknown"}, Sel: &ast.Ident{Name: "Use"}},
	}
	tryUseCall(call, map[string]*routerInfo{})
}
