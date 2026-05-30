//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberContextType_WrongNames 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberContextType_WrongNames(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "http"},
			Sel: &ast.Ident{Name: "Request"},
		},
	}
	if isFiberContextType(expr) {
		t.Fatal("expected false for *http.Request")
	}
}
