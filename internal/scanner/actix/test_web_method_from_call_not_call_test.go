//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestWebMethodFromCall_NotCall 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWebMethodFromCall_NotCall(t *testing.T) {

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
	if got := webMethodFromCall(id, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
