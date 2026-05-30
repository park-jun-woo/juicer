//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractDeclaratorLiteral 테스트
package spring

import "testing"

func TestExtractDeclaratorLiteral(t *testing.T) {
	root, src := parseS(t, `class C { static final String X = "hello"; }`)
	decls := findAllByType(root, "variable_declarator")
	if len(decls) == 0 {
		t.Skip("no declarator")
	}
	if got := extractDeclaratorLiteral(decls[0], src); got != "hello" {
		t.Fatalf("got %q", got)
	}
}
