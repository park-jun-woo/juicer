//ff:func feature=scan type=test control=sequence topic=spring
//ff:what sParam 테스트 헬퍼
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func sParam(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	root, b := sParse(t, src)
	return sFirst(t, root, "formal_parameter"), b
}
