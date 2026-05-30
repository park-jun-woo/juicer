//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestWalkNodes_Round5 테스트
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkNodes_Round5(t *testing.T) {
	root, _ := parseCS(t, "class C { void M() {} }")
	count := 0
	walkNodes(root, func(n *sitter.Node) { count++ })
	if count < 3 {
		t.Fatalf("walkNodes visited too few nodes: %d", count)
	}
}
