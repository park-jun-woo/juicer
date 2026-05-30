//ff:func feature=scan type=test control=sequence topic=flask
//ff:what firstAssignment 테스트 헬퍼
package flask

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstAssignment(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	assigns := findAllByType(root, "assignment")
	if len(assigns) == 0 {
		t.Fatal("no assignment")
	}
	return assigns[0], b
}
