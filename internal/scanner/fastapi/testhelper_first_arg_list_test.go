//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what firstArgList 테스트 헬퍼
package fastapi

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstArgList(t *testing.T, src []byte) (*sitter.Node, []byte) {
	t.Helper()
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	al := findAllByType(root, "argument_list")
	if len(al) == 0 {
		t.Fatal("no argument_list")
	}
	return al[0], src
}
