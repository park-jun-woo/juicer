//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestWalkMethodChain_NonCallNode 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkMethodChain_NonCallNode(t *testing.T) {

	src := []byte(`fn f() { lone; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	var id *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if id != nil {
			return
		}
		if n.Type() == "identifier" && nodeText(n, src) == "lone" {
			id = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	count := 0
	walkMethodChain(id, src, "service", func(args *sitter.Node) { count++ })
	if count != 0 {
		t.Fatalf("expected 0 callbacks, got %d", count)
	}
}
