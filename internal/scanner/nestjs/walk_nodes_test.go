//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what walkNodes 테스트
package nestjs

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestWalkNodes_Count(t *testing.T) {
	src := []byte(`const x = 1;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	walkNodes(root, func(n *sitter.Node) {
		count++
	})
	if count == 0 {
		t.Fatal("expected at least one node")
	}
}
