//ff:func feature=scan type=test control=sequence
//ff:what TestIsEchoInit_New 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestIsEchoInit_New(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "echo"}, Sel: &ast.Ident{Name: "New"}}
	if !isEchoInit(sel, "echo") {
		t.Fatal("expected true")
	}
}
