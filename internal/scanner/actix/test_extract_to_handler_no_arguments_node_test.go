//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractToHandler_NoArgumentsNode 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractToHandler_NoArgumentsNode(t *testing.T) {

	src := []byte(`fn f() { lone; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	var ident *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if ident != nil {
			return
		}
		if n.Type() == "identifier" && nodeText(n, src) == "lone" {
			ident = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if ident == nil {
		t.Fatal("identifier not found")
	}
	if got := extractToHandler(ident, src); got != "" {
		t.Fatalf("expected empty for node without arguments, got %q", got)
	}
}
