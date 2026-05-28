//ff:func feature=scan type=test control=sequence
//ff:what TestIsEchoInit_NotNew 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestIsEchoInit_NotNew(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "echo"}, Sel: &ast.Ident{Name: "Default"}}
	if isEchoInit(sel, "echo") {
		t.Fatal("expected false — Echo only has New(), not Default()")
	}
}
