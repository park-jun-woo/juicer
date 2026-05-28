//ff:func feature=scan type=test control=sequence
//ff:what TestIsEchoRouterType_Wrong 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestIsEchoRouterType_Wrong(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "echo"},
			Sel: &ast.Ident{Name: "Other"},
		},
	}
	if isEchoRouterType(expr, "echo") {
		t.Fatal("expected false for *echo.Other")
	}
}
