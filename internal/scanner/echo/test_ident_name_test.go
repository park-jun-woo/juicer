//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestIdentName 테스트
package echo

import "testing"

func TestIdentName(t *testing.T) {
	if identName(parseExpr(t, "foo")) != "foo" {
		t.Fatal("ident")
	}
	if identName(parseExpr(t, "a.b")) != "" {
		t.Fatal("selector should be empty")
	}
}
