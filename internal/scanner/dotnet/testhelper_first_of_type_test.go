//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what firstOfType 테스트 헬퍼
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

// firstOfType returns the first descendant node of the given type.
func firstOfType(t *testing.T, root *sitter.Node, typ string) *sitter.Node {
	t.Helper()
	nodes := findAllByType(root, typ)
	if len(nodes) == 0 {
		t.Fatalf("no %s node found", typ)
	}
	return nodes[0]
}
