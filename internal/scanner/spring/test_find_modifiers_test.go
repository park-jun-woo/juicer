//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestFindModifiers 테스트
package spring

import "testing"

func TestFindModifiers(t *testing.T) {
	root, _ := parseS(t, `class C { public int x; }`)
	field := findAllByType(root, "field_declaration")[0]
	if findModifiers(field) == nil {
		t.Fatal("modifiers")
	}
}
