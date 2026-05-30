//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestIsEchoInit 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestIsEchoInit(t *testing.T) {
	sel := parseExpr(t, "e.New").(*ast.SelectorExpr)
	if !isEchoInit(sel, "e") {
		t.Fatal("expected true for e.New")
	}
	if isEchoInit(sel, "other") {
		t.Fatal("wrong alias should be false")
	}
	sel2 := parseExpr(t, "e.Start").(*ast.SelectorExpr)
	if isEchoInit(sel2, "e") {
		t.Fatal("e.Start should be false")
	}
}
