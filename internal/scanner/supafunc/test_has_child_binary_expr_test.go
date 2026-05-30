//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestHasChildBinaryExpr 테스트
package supafunc

import "testing"

func TestHasChildBinaryExpr(t *testing.T) {
	fi := mustParse(t, []byte(`if (a === b) {}`))
	parens := findAllByType(fi.Root, "parenthesized_expression")
	if len(parens) == 0 {
		t.Skip("no parenthesized expr")
	}
	if !hasChildBinaryExpr(parens[0]) {
		t.Fatal("expected binary child")
	}
	fi2 := mustParse(t, []byte(`if (a) {}`))
	parens2 := findAllByType(fi2.Root, "parenthesized_expression")
	if len(parens2) > 0 && hasChildBinaryExpr(parens2[0]) {
		t.Fatal("no binary child expected")
	}
}
