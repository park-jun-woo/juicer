//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what firstFunc 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstFunc(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	for _, typ := range []string{"function_declaration", "function", "arrow_function"} {
		if ns := findAllByType(fi.Root, typ); len(ns) > 0 {
			return ns[0], fi.Src
		}
	}
	t.Fatalf("no function in %q", src)
	return nil, nil
}
