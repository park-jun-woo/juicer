//ff:func feature=scan type=test control=sequence
//ff:what TestIsEchoContextType_NotEcho 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestIsEchoContextType_NotEcho(t *testing.T) {
	expr := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "gin"},
		Sel: &ast.Ident{Name: "Context"},
	}
	if isEchoContextType(expr) {
		t.Fatal("expected false — gin.Context is not echo.Context")
	}
}
