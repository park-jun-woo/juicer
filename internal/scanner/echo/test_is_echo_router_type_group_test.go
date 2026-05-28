//ff:func feature=scan type=test control=sequence
//ff:what TestIsEchoRouterType_Group 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestIsEchoRouterType_Group(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "echo"},
			Sel: &ast.Ident{Name: "Group"},
		},
	}
	if !isEchoRouterType(expr, "echo") {
		t.Fatal("expected true for *echo.Group")
	}
}
