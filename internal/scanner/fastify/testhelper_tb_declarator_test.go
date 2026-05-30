//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what tbDeclarator 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func tbDeclarator(t *testing.T, src string) (*sitter.Node, *fileInfo) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	ds := findAllByType(fi.Root, "variable_declarator")
	if len(ds) == 0 {
		t.Fatalf("no declarator in %q", src)
	}
	return ds[0], fi
}
