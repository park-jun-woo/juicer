//ff:func feature=scan type=test control=sequence topic=spring
//ff:what parseS 테스트 헬퍼
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func parseS(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	return root, b
}
