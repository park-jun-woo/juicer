//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what firstFieldDecl 테스트 헬퍼
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstFieldDecl(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	src := []byte(javaSrc)
	root, err := parseJava(src)
	if err != nil {
		t.Fatal(err)
	}
	fields := findAllByType(root, "field_declaration")
	if len(fields) == 0 {
		t.Fatal("no field_declaration")
	}
	return fields[0], src
}
