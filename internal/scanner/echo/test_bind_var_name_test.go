//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestBindVarName 테스트
package echo

import "testing"

func TestBindVarName(t *testing.T) {
	if bindVarName(parseExpr(t, "&dto")) != "dto" {
		t.Fatal("address-of")
	}
	if bindVarName(parseExpr(t, "dto")) != "dto" {
		t.Fatal("plain")
	}
}
