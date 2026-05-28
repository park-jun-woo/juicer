//ff:func feature=scan type=test control=sequence
//ff:what TestIsEchoContextType_Valid 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestIsEchoContextType_Valid(t *testing.T) {
	// Echo uses echo.Context (interface, not pointer)
	expr := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "echo"},
		Sel: &ast.Ident{Name: "Context"},
	}
	if !isEchoContextType(expr) {
		t.Fatal("expected true")
	}
}
