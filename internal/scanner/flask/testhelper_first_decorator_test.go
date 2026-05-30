//ff:func feature=scan type=test control=sequence topic=flask
//ff:what firstDecorator 테스트 헬퍼
package flask

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstDecorator(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Fatal("no decorator")
	}
	return decs[0], b
}
