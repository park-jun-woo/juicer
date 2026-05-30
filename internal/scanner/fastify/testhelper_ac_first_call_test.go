//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what acFirstCall 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func acFirstCall(t *testing.T, src string) (*fileInfo, []*sitter.Node) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	return fi, findAllByType(fi.Root, "call_expression")
}
