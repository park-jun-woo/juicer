//ff:func feature=scan type=test control=sequence topic=zod
//ff:what findFirstCall 테스트 헬퍼
package zod

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

// findFirstCall returns the first call_expression node in the tree.
func findFirstCall(t *testing.T, root *sitter.Node) *sitter.Node {
	t.Helper()
	var found *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if found == nil && n.Type() == "call_expression" {
			found = n
		}
	})
	if found == nil {
		t.Fatal("no call_expression found")
	}
	return found
}
