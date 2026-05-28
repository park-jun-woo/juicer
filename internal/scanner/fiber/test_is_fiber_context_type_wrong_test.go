//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberContextType_Wrong 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberContextType_Wrong(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "gin"},
			Sel: &ast.Ident{Name: "Context"},
		},
	}
	if isFiberContextType(expr) {
		t.Fatal("expected false — this is gin.Context, not fiber.Ctx")
	}
}
