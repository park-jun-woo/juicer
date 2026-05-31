//ff:func feature=scan type=test topic=joi control=sequence
//ff:what walkNodes 전위 순회로 모든 노드 방문 테스트
package joi

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestWalkNodes(t *testing.T) {
	root, _ := parseJoiTS(t, `const a = 1;`)
	count := 0
	sawProgram := false
	walkNodes(root, func(n *sitter.Node) {
		count++
		if n.Type() == "program" {
			sawProgram = true
		}
	})
	if count < 3 {
		t.Errorf("expected several nodes, got %d", count)
	}
	if !sawProgram {
		t.Error("root program node should be visited")
	}
}
