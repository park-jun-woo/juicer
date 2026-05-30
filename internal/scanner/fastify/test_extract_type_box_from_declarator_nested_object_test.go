//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractTypeBoxFromDeclarator_NestedObject 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractTypeBoxFromDeclarator_NestedObject(t *testing.T) {
	d, fi := tbDeclarator(t, "const S = Type.Object({ inner: Type.Object({ x: Type.Number() }) });\n")
	vars := map[string]*sitter.Node{}
	extractTypeBoxFromDeclarator(d, fi, vars)
	if vars["S"] == nil {
		t.Fatalf("expected S captured, got %v", vars)
	}
}
