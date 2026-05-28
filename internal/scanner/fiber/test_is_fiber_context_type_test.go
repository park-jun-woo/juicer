//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberContextType_Valid 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberContextType_Valid(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "fiber"},
			Sel: &ast.Ident{Name: "Ctx"},
		},
	}
	if !isFiberContextType(expr) {
		t.Fatal("expected true")
	}
}
