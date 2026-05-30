//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what firstAssignment 테스트 헬퍼
package fastapi

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstAssignment(t *testing.T, src []byte) (*sitter.Node, []byte) {
	t.Helper()
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	a := findAllByType(root, "assignment")
	if len(a) == 0 {
		t.Fatal("no assignment")
	}
	return a[0], src
}
