//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindFieldType_TypeIdentifier 테스트
package actix

import "testing"

func TestFindFieldType_TypeIdentifier(t *testing.T) {
	src := []byte(`struct S { name: String }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fd := firstFieldDecl(root)
	if fd == nil {
		t.Fatal("no field_declaration")
	}
	ty := findFieldType(fd)
	if ty == nil {
		t.Fatal("expected a type node")
	}
	if nodeText(ty, src) != "String" {
		t.Errorf("type = %q, want String", nodeText(ty, src))
	}
}
