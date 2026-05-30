//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what firstParam 테스트 헬퍼
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstParam(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(javaSrc)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	params := findAllByType(root, "formal_parameter")
	if len(params) == 0 {
		t.Fatal("no formal_parameter")
	}
	return params[0], b
}
