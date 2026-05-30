//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractTypeBoxFromDeclarator_Match 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractTypeBoxFromDeclarator_Match(t *testing.T) {
	d, fi := tbDeclarator(t, "const S = Type.Object({ a: Type.String() });\n")
	vars := map[string]*sitter.Node{}
	extractTypeBoxFromDeclarator(d, fi, vars)
	if _, ok := vars["S"]; !ok {
		t.Fatalf("expected S captured, got %v", vars)
	}
}
