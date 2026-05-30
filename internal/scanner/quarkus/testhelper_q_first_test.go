//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what qFirst 테스트 헬퍼
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func qFirst(t *testing.T, root *sitter.Node, typ string) *sitter.Node {
	t.Helper()
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == typ {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if found == nil {
		t.Fatalf("no %s node", typ)
	}
	return found
}
