//ff:func feature=scan type=test control=sequence topic=spring
//ff:what firstMethodS 테스트 헬퍼
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstMethodS(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	root, src := parseS(t, javaSrc)
	methods := findAllByType(root, "method_declaration")
	if len(methods) == 0 {
		t.Fatal("no method")
	}
	return methods[0], src
}
