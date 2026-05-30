//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what parseCS 테스트 헬퍼
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func parseCS(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parseCSharp(b)
	if err != nil {
		t.Fatal(err)
	}
	return root, b
}
