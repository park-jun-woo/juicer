//ff:func feature=scan type=test control=sequence topic=spring
//ff:what firstParamS 테스트 헬퍼
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstParamS(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	root, src := parseS(t, javaSrc)
	params := findAllByType(root, "formal_parameter")
	if len(params) == 0 {
		t.Fatal("no param")
	}
	return params[0], src
}
