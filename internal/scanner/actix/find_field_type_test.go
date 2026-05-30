//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findFieldType — field_declaration의 타입 노드 탐색을 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstFieldDecl(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "field_declaration" {
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

func TestFindFieldType_NotFound(t *testing.T) {
	// An empty block has no recognized type child -> nil.
	src := []byte(`fn f() {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	var block *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if block != nil {
			return
		}
		if n.Type() == "block" {
			block = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if block == nil {
		t.Fatal("no block")
	}
	if ty := findFieldType(block); ty != nil {
		t.Fatalf("expected nil, got %v (%s)", ty, ty.Type())
	}
}
