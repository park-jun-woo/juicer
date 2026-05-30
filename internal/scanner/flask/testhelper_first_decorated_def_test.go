//ff:func feature=scan type=test control=sequence topic=flask
//ff:what firstDecoratedDef 테스트 헬퍼
package flask

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstDecoratedDef(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	defs := findAllByType(root, "decorated_definition")
	if len(defs) == 0 {
		t.Fatal("no decorated_definition")
	}
	return defs[0], b
}
