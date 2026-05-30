//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what firstCall 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstCall(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatalf("no call in %q", src)
	}
	return calls[0], fi.Src
}
