//ff:func feature=scan type=test control=sequence topic=flask
//ff:what argListOf 테스트 헬퍼
package flask

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func argListOf(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	args := findAllByType(root, "argument_list")
	if len(args) == 0 {
		t.Fatal("no argument_list")
	}
	return args[0], b
}
