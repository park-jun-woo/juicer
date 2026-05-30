//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestMatchStatusArgChild_IntLiteral 테스트
package spring

import "testing"

func TestMatchStatusArgChild_IntLiteral(t *testing.T) {
	root, src := parseS(t, `class C { void m() { status(404); } }`)
	lits := findAllByType(root, "decimal_integer_literal")
	if len(lits) == 0 {
		t.Skip("no literal")
	}
	if got := matchStatusArgChild(lits[0], src); got != "404" {
		t.Fatalf("got %q", got)
	}
}
