//ff:func feature=scan type=test control=sequence
//ff:what TestIsEchoInit_WrongAlias 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestIsEchoInit_WrongAlias(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "other"}, Sel: &ast.Ident{Name: "New"}}
	if isEchoInit(sel, "echo") {
		t.Fatal("expected false")
	}
}
