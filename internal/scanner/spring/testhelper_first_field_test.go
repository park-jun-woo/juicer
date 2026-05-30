//ff:func feature=scan type=test control=sequence topic=spring
//ff:what firstField 테스트 헬퍼
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstField(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	root, src := parseS(t, javaSrc)
	fields := findAllByType(root, "field_declaration")
	if len(fields) == 0 {
		t.Fatal("no field")
	}
	return fields[0], src
}
