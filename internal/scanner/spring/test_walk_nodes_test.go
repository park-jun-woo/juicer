//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestWalkNodes 테스트
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkNodes(t *testing.T) {
	root, _ := parseS(t, `class C { void a() {} void b() {} }`)
	count := 0
	walkNodes(root, func(n *sitter.Node) {
		if n.Type() == "method_declaration" {
			count++
		}
	})
	if count != 2 {
		t.Fatalf("got %d", count)
	}
}
