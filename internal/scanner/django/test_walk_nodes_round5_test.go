//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestWalkNodes_Round5 테스트
package django

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestWalkNodes_Round5(t *testing.T) {
	root, _ := parsePython([]byte("x = 1\ny = 2\n"))
	n := 0
	walkNodes(root, func(*sitter.Node) { n++ })
	if n < 3 {
		t.Fatalf("too few nodes: %d", n)
	}
}
