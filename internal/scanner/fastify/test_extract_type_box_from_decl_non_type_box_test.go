//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractTypeBoxFromDecl_NonTypeBox 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractTypeBoxFromDecl_NonTypeBox(t *testing.T) {
	fi := mustParse(t, []byte("const x = foo();\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	vars := map[string]*sitter.Node{}
	for _, d := range decls {
		extractTypeBoxFromDecl(d, fi, vars)
	}
	if len(vars) != 0 {
		t.Fatalf("expected no vars, got %v", vars)
	}
}
