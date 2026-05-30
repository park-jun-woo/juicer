//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findAllByType — 지정 타입 노드 수집을 검증
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

func TestFindAllByType_NoMatch(t *testing.T) {
	src := []byte(`fn f() {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if nodes := findAllByType(root, "struct_item"); len(nodes) != 0 {
		t.Fatalf("expected no matches, got %d", len(nodes))
	}
}
