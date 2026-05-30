//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what firstParamCS 테스트 헬퍼
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstParamCS(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	root, b := parseCS(t, src)
	params := findAllByType(root, "parameter")
	if len(params) == 0 {
		t.Fatal("no param")
	}
	return params[0], b
}
