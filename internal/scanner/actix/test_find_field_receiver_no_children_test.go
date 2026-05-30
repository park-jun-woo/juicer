//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindFieldReceiver_NoChildren 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestFindFieldReceiver_NoChildren(t *testing.T) {

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
