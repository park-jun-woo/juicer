//ff:func feature=scan type=test control=sequence topic=flask
//ff:what firstCall 테스트 헬퍼
package flask

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstCall(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	calls := findAllByType(root, "call")
	if len(calls) == 0 {
		t.Fatal("no call")
	}
	return calls[0], b
}
