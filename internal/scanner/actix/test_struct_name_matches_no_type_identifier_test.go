//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestStructNameMatches_NoTypeIdentifier 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestStructNameMatches_NoTypeIdentifier(t *testing.T) {

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
