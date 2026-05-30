//ff:func feature=scan type=test control=sequence topic=hono
//ff:what firstCallExpr 테스트 헬퍼
package hono

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstCallExpr(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call_expression")
	}
	return calls[0], fi.Src
}
