//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestFindModifiers 테스트
package quarkus

import "testing"

func TestFindModifiers(t *testing.T) {
	root, _ := parseJava([]byte(`class C { public int x; }`))
	fields := findAllByType(root, "field_declaration")
	if findModifiers(fields[0]) == nil {
		t.Fatal("expected modifiers")
	}
}
