//ff:func feature=scan type=test control=sequence topic=actix
//ff:what childrenOfType — 직접 자식 중 지정 타입 수집을 검증
package actix

import "testing"

func TestChildrenOfType(t *testing.T) {
	src := []byte(`struct S { a: i32, b: i32 }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	sn := firstStructNode(root)
	if sn == nil {
		t.Fatal("no struct")
	}
	fdl := findChildByType(sn, "field_declaration_list")
	if fdl == nil {
		t.Fatal("no field_declaration_list")
	}
	fields := childrenOfType(fdl, "field_declaration")
	if len(fields) != 2 {
		t.Fatalf("expected 2 field_declaration children, got %d", len(fields))
	}
	// No matching type -> empty.
	if got := childrenOfType(fdl, "struct_item"); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}
