//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what firstClass 테스트 헬퍼
package fastapi

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstClass(t *testing.T, src []byte) (*sitter.Node, []byte) {
	t.Helper()
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	cls := findAllByType(root, "class_definition")
	if len(cls) == 0 {
		t.Fatal("no class_definition")
	}
	return cls[0], src
}
