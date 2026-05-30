//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestIsStaticField 테스트
package quarkus

import "testing"

func TestIsStaticField(t *testing.T) {
	root, _ := parseJava([]byte(`class C { static int a; int b; }`))
	fields := findAllByType(root, "field_declaration")
	if !isStaticField(fields[0]) {
		t.Fatal("a is static")
	}
	if isStaticField(fields[1]) {
		t.Fatal("b is not static")
	}
}
