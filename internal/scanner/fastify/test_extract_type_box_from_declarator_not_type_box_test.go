//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractTypeBoxFromDeclarator_NotTypeBox 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractTypeBoxFromDeclarator_NotTypeBox(t *testing.T) {
	d, fi := tbDeclarator(t, "const x = Other.Object({});\n")
	vars := map[string]*sitter.Node{}
	extractTypeBoxFromDeclarator(d, fi, vars)
	if len(vars) != 0 {
		t.Fatalf("expected no vars for non-TypeBox, got %v", vars)
	}
}
