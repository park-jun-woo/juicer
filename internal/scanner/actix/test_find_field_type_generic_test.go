//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindFieldType_Generic 테스트
package actix

import "testing"

func TestFindFieldType_Generic(t *testing.T) {
	src := []byte(`struct S { id: Option<i64> }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fd := firstFieldDecl(root)
	if fd == nil {
		t.Fatal("no field_declaration")
	}
	ty := findFieldType(fd)
	if ty == nil || ty.Type() != "generic_type" {
		t.Fatalf("expected generic_type, got %v", ty)
	}
}
