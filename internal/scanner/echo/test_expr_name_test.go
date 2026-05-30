//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestExprName 테스트
package echo

import "testing"

func TestExprName(t *testing.T) {
	if exprName(parseExpr(t, "foo")) != "foo" {
		t.Fatal("ident")
	}
	if exprName(parseExpr(t, "pkg.Fn")) != "pkg.Fn" {
		t.Fatal("selector")
	}
	if exprName(parseExpr(t, "fn()")) != "fn()" {
		t.Fatal("call")
	}
	if exprName(parseExpr(t, "func(){}")) != "(inline)" {
		t.Fatal("funclit")
	}
}
