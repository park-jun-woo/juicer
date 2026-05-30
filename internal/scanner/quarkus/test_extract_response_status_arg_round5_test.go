//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractResponseStatusArg_Round5 테스트
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractResponseStatusArg_Round5(t *testing.T) {
	root, src := qParse(t, `class C { void m() { Response.status(201).build(); } }`)
	// find the status(201) invocation
	var inv *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if inv != nil {
			return
		}
		if n.Type() == "method_invocation" {
			name := n.ChildByFieldName("name")
			if name != nil && nodeText(name, src) == "status" {
				inv = n
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if inv == nil {
		t.Fatal("no status invocation")
	}
	got := extractResponseStatusArg(inv, src)
	if got != "201" {
		t.Fatalf("status: %q", got)
	}
}
