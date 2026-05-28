//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterType_Wrong 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberRouterType_Wrong(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "fiber"},
			Sel: &ast.Ident{Name: "Config"},
		},
	}
	if isFiberRouterType(expr, "fiber") {
		t.Fatal("expected false — Config is not a router type")
	}
}
