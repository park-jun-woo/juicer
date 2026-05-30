//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterType_NoStarGroup 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberRouterType_NoStarGroup(t *testing.T) {

	expr := &ast.SelectorExpr{X: &ast.Ident{Name: "fiber"}, Sel: &ast.Ident{Name: "Router"}}
	_ = isFiberRouterType(expr, "fiber")
}
