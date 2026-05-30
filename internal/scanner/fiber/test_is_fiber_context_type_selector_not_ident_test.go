//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberContextType_SelectorNotIdent 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberContextType_SelectorNotIdent(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.CallExpr{Fun: &ast.Ident{Name: "f"}},
			Sel: &ast.Ident{Name: "Ctx"},
		},
	}
	if isFiberContextType(expr) {
		t.Fatal("expected false — selector X not an ident")
	}
}
