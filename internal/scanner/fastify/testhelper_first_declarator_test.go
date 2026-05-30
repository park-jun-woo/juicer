//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what firstDeclarator 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstDeclarator(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	ds := findAllByType(fi.Root, "variable_declarator")
	if len(ds) == 0 {
		t.Fatalf("no variable_declarator in %q", src)
	}
	return ds[0], fi.Src
}
