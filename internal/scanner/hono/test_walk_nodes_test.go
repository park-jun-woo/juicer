//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestWalkNodes 테스트
package hono

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkNodes(t *testing.T) {
	fi := mustParse(t, []byte(`a(b());`+"\n"))
	count := 0
	walkNodes(fi.Root, func(n *sitter.Node) {
		if n.Type() == "call_expression" {
			count++
		}
	})
	if count != 2 {
		t.Fatalf("expected 2 call_expression visits, got %d", count)
	}
}
