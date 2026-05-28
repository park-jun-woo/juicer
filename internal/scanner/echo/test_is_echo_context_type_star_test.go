//ff:func feature=scan type=test control=sequence
//ff:what TestIsEchoContextType_StarExpr 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestIsEchoContextType_StarExpr(t *testing.T) {
	// *echo.Context — not the typical Echo pattern (echo.Context is an interface)
	// but isEchoContextType checks for non-pointer form
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "echo"},
			Sel: &ast.Ident{Name: "Context"},
		},
	}
	if isEchoContextType(expr) {
		t.Fatal("expected false — Echo uses interface, not pointer")
	}
}
