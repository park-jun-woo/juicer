//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what firstNodeOfType 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstNodeOfType(t *testing.T, src, typ string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	ns := findAllByType(fi.Root, typ)
	if len(ns) == 0 {
		t.Fatalf("no %s node in %q", typ, src)
	}
	return ns[0], fi.Src
}
