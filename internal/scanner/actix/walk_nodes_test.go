//ff:func feature=scan type=test control=sequence topic=actix
//ff:what walkNodes — 서브트리 전체 순회를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestWalkNodes(t *testing.T) {
	src := []byte(`fn f() { a; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	visited := 0
	sawRoot := false
	walkNodes(root, func(n *sitter.Node) {
		visited++
		if n == root {
			sawRoot = true
		}
	})
	if !sawRoot {
		t.Error("expected root to be visited")
	}
	// A non-trivial tree visits more than just the root.
	if visited < 2 {
		t.Errorf("expected multiple nodes visited, got %d", visited)
	}
}
