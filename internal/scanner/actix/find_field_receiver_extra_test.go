//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findFieldReceiver — 자식 없는 노드는 nil 반환을 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestFindFieldReceiver_HasReceiver(t *testing.T) {
	src := []byte(`fn f() { foo.bar; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fe := firstFieldExpr(root)
	if fe == nil {
		t.Fatal("no field_expression")
	}
	recv := findFieldReceiver(fe)
	if recv == nil {
		t.Fatal("expected non-nil receiver")
	}
	if nodeText(recv, src) != "foo" {
		t.Errorf("receiver = %q, want foo", nodeText(recv, src))
	}
}

func TestFindFieldReceiver_NoChildren(t *testing.T) {
	// A leaf identifier node has no children -> nil.
	src := []byte(`fn f() { lone; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	var leaf *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if leaf != nil {
			return
		}
		if n.Type() == "identifier" && nodeText(n, src) == "lone" {
			leaf = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if leaf == nil {
		t.Fatal("identifier not found")
	}
	if leaf.ChildCount() != 0 {
		t.Fatalf("expected leaf node, got %d children", leaf.ChildCount())
	}
	if recv := findFieldReceiver(leaf); recv != nil {
		t.Fatalf("expected nil receiver, got %v", recv)
	}
}
