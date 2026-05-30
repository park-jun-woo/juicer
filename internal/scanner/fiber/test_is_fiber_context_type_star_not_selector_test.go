//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberContextType_StarNotSelector 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberContextType_StarNotSelector(t *testing.T) {
	expr := &ast.StarExpr{X: &ast.Ident{Name: "Ctx"}}
	if isFiberContextType(expr) {
		t.Fatal("expected false — *Ctx without selector")
	}
}
