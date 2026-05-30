//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseOneField_NoName 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestParseOneField_NoName(t *testing.T) {

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
	if f := parseOneField(block, src, nil); f != nil {
		t.Fatalf("expected nil for node without field_identifier, got %+v", f)
	}
}
