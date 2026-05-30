//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberContextType_NotStar 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberContextType_NotStar(t *testing.T) {
	expr := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "fiber"},
		Sel: &ast.Ident{Name: "Ctx"},
	}
	if isFiberContextType(expr) {
		t.Fatal("expected false — must be *fiber.Ctx with star")
	}
}
