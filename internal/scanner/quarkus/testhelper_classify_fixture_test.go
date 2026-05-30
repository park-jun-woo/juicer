//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what classifyFixture 테스트 헬퍼
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func classifyFixture(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	root, b := qParse(t, src)
	param := qFirst(t, root, "formal_parameter")
	return param, b
}
