//ff:func feature=scan type=test control=sequence topic=zod
//ff:what walkNodes 테스트 (round5)
package zod

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestWalkNodes_Round5(t *testing.T) {
	root, _ := parseTS(t, `const a = 1; const b = 2;`)
	count := 0
	walkNodes(root, func(n *sitter.Node) { count++ })
	if count < 3 {
		t.Fatalf("walkNodes visited too few nodes: %d", count)
	}
}
