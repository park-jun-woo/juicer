//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestWalkNodes 테스트
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkNodes(t *testing.T) {
	root, _ := parseJava([]byte(`class C { void a() {} void b() {} }`))
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
