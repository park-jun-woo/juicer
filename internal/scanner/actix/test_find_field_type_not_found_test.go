//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindFieldType_NotFound 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestFindFieldType_NotFound(t *testing.T) {

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
