//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what walkNodes 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestWalkNodes(t *testing.T) {
	fi := mustParse(t, []byte("const x = 1;\n"))
	count := 0
	identifiers := 0
	walkNodes(fi.Root, func(n *sitter.Node) {
		count++
		if n.Type() == "identifier" {
			identifiers++
		}
	})
	if count < 3 {
		t.Fatalf("expected to visit several nodes, got %d", count)
	}
	if identifiers < 1 {
		t.Fatalf("expected at least one identifier visited, got %d", identifiers)
	}
}
