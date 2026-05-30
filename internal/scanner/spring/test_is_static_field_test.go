//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestIsStaticField 테스트
package spring

import "testing"

func TestIsStaticField(t *testing.T) {
	root, _ := parseS(t, `class C { static int a; int b; }`)
	fields := findAllByType(root, "field_declaration")
	if !isStaticField(fields[0]) || isStaticField(fields[1]) {
		t.Fatal("static")
	}
}
