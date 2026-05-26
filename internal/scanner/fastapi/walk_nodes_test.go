//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what walkNodes 테스트
package fastapi

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestWalkNodes(t *testing.T) {
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	walkNodes(root, func(n *sitter.Node) {
		count++
	})
	if count < 3 {
		t.Fatalf("expected at least 3 nodes, got %d", count)
	}
}
