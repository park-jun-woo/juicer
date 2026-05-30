//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestExprString 테스트
package echo

import "testing"

func TestExprString(t *testing.T) {
	if exprString(parseExpr(t, "foo")) != "foo" {
		t.Fatal("ident")
	}
	if exprString(parseExpr(t, "pkg.Type")) != "pkg.Type" {
		t.Fatal("selector")
	}
	if exprString(parseExpr(t, "T{}")) != "T{}" {
		t.Fatal("composite")
	}
}
