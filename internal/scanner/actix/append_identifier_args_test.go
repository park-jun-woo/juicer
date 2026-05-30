//ff:func feature=scan type=test control=sequence topic=actix
//ff:what appendIdentifierArgs — arguments 노드의 식별자 인자만 수집함을 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestAppendIdentifierArgs(t *testing.T) {
	// "plain" and "other" are identifiers; the string literal and the scoped
	// path are not "identifier"-typed children, so they are skipped.
	src := []byte(`
fn f() {
    g(plain, "lit", view::about, other);
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}

	var got []string
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "arguments" {
			got = appendIdentifierArgs(n, src, got)
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)

	want := map[string]bool{"plain": true, "other": true}
	for _, g := range got {
		if !want[g] {
			t.Errorf("unexpected identifier collected: %q (got %v)", g, got)
		}
	}
	if len(got) != 2 {
		t.Fatalf("expected exactly 2 identifier args, got %v", got)
	}
}
