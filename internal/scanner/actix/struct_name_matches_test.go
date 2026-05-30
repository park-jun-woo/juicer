//ff:func feature=scan type=test control=sequence topic=actix
//ff:what structNameMatches — struct 타입명 일치 판별 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestStructNameMatches(t *testing.T) {
	src := []byte(`struct User { id: i64 }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	sn := firstStructNode(root)
	if sn == nil {
		t.Fatal("no struct")
	}
	if !structNameMatches(sn, src, "User") {
		t.Error("expected match for User")
	}
	if structNameMatches(sn, src, "Other") {
		t.Error("expected no match for Other")
	}
}

func TestStructNameMatches_NoTypeIdentifier(t *testing.T) {
	// A block node has no type_identifier -> false.
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
	if structNameMatches(block, src, "User") {
		t.Error("expected false for node without type_identifier")
	}
}
