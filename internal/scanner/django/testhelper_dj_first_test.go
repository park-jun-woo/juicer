//ff:func feature=scan type=test control=sequence topic=django
//ff:what djFirst 테스트 헬퍼
package django

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func djFirst(t *testing.T, root *sitter.Node, typ string) *sitter.Node {
	t.Helper()
	var found *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if found == nil && n.Type() == typ {
			found = n
		}
	})
	if found == nil {
		t.Fatalf("no %s node", typ)
	}
	return found
}
