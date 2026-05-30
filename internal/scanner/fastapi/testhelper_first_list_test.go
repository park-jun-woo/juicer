//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what firstList 테스트 헬퍼
package fastapi

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstList(t *testing.T, src []byte) (*sitter.Node, []byte) {
	t.Helper()
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	lists := findAllByType(root, "list")
	if len(lists) == 0 {
		t.Fatal("no list")
	}
	return lists[0], src
}
