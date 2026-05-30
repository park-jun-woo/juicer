//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestWalkNodes 테스트
package supafunc

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkNodes(t *testing.T) {
	fi := mustParse(t, []byte(`a(b());`))
	count := 0
	walkNodes(fi.Root, func(n *sitter.Node) {
		if n.Type() == "call_expression" {
			count++
		}
	})
	if count != 2 {
		t.Fatalf("got %d", count)
	}
}
