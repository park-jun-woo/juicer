//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what importClause 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func importClause(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	cs := findAllByType(fi.Root, "import_clause")
	if len(cs) == 0 {
		t.Fatalf("no import_clause in %q", src)
	}
	return cs[0], fi.Src
}
