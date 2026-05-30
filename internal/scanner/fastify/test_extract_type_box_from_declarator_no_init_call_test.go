//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractTypeBoxFromDeclarator_NoInitCall 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractTypeBoxFromDeclarator_NoInitCall(t *testing.T) {
	d, fi := tbDeclarator(t, "const x = 5;\n")
	vars := map[string]*sitter.Node{}
	extractTypeBoxFromDeclarator(d, fi, vars)
	if len(vars) != 0 {
		t.Fatalf("expected no vars, got %v", vars)
	}
}
