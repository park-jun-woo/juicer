//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestHasModifiers 테스트
package spring

import "testing"

func TestHasModifiers(t *testing.T) {
	root, src := parseS(t, `class C { public static final int X = 1; }`)
	field := findAllByType(root, "field_declaration")[0]
	if !hasModifiers(field, src, "static", "final") {
		t.Fatal("expected static final")
	}
	if hasModifiers(field, src, "private") {
		t.Fatal("not private")
	}
}
