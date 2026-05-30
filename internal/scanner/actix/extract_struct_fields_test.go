//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractStructFields — struct 필드 파싱 및 필드없는 struct nil 분기 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstStructNode(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "struct_item" {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}

func TestExtractStructFields(t *testing.T) {
	src := []byte(`
struct User {
    id: i64,
    name: String,
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	sn := firstStructNode(root)
	if sn == nil {
		t.Fatal("no struct found")
	}
	fields := extractStructFields(sn, src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d: %+v", len(fields), fields)
	}
	if fields[0].Name != "id" || fields[1].Name != "name" {
		t.Errorf("unexpected fields: %+v", fields)
	}
}

func TestExtractStructFields_Unit(t *testing.T) {
	// Unit struct has no field_declaration_list.
	src := []byte(`struct Marker;`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	sn := firstStructNode(root)
	if sn == nil {
		t.Fatal("no struct found")
	}
	if fields := extractStructFields(sn, src); fields != nil {
		t.Fatalf("expected nil for unit struct, got %+v", fields)
	}
}
