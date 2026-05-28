//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterType_App 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberRouterType_App(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "fiber"},
			Sel: &ast.Ident{Name: "App"},
		},
	}
	if !isFiberRouterType(expr, "fiber") {
		t.Fatal("expected true")
	}
}
