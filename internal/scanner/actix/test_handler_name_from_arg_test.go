//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestHandlerNameFromArg — identifier/scoped/generic/closure 핸들러명 해석을 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestHandlerNameFromArg(t *testing.T) {
	src := []byte(`
fn f() {
    g(plain);
    g(view::about);
    g(api::search::<String>);
    g(|| async { 1 });
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
			for i := 0; i < int(n.ChildCount()); i++ {
				if name := handlerNameFromArg(n.Child(i), src); name != "" {
					got = append(got, name)
				}
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)

	want := map[string]bool{"plain": false, "about": false, "search": false}
	for _, g := range got {
		if _, ok := want[g]; ok {
			want[g] = true
		}
	}
	for k, seen := range want {
		if !seen {
			t.Errorf("expected handler name %q resolved, got %v", k, got)
		}
	}
	// closure resolves to "" (not in got)
	for _, g := range got {
		if g == "" {
			t.Error("closure must resolve to empty name")
		}
	}
}
