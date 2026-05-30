//ff:func feature=scan type=test control=sequence topic=flask
//ff:what walkNodes 테스트
package flask

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestWalkNodes(t *testing.T) {
	b := []byte("x = 1\n")
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	idents := 0
	walkNodes(root, func(n *sitter.Node) {
		count++
		if n.Type() == "identifier" {
			idents++
		}
	})
	if count < 3 {
		t.Fatalf("expected several nodes, got %d", count)
	}
	if idents < 1 {
		t.Fatalf("expected at least 1 identifier, got %d", idents)
	}
}
