//ff:func feature=scan type=test topic=flask control=sequence
//ff:what flask 테스트 헬퍼 — Python 소스 파싱 후 첫 function_definition 반환
package flask

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func flaskFuncDef(t *testing.T, code string) (*sitter.Node, []byte) {
	t.Helper()
	src := []byte(code)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	defs := findAllByType(root, "function_definition")
	if len(defs) == 0 {
		t.Fatal("no function_definition")
	}
	return defs[0], src
}
