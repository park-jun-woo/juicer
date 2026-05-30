//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindAllByType 테스트
package actix

import "testing"

func TestFindAllByType(t *testing.T) {
	src := []byte(`
struct A { x: i32 }
struct B { y: i32 }
fn f() {}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	structs := findAllByType(root, "struct_item")
	if len(structs) != 2 {
		t.Fatalf("expected 2 struct_item nodes, got %d", len(structs))
	}
}
