//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractTypeBoxFromDeclarator 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
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

func TestExtractTypeBoxFromDeclarator_Match(t *testing.T) {
	d, fi := tbDeclarator(t, "const S = Type.Object({ a: Type.String() });\n")
	vars := map[string]*sitter.Node{}
	extractTypeBoxFromDeclarator(d, fi, vars)
	if _, ok := vars["S"]; !ok {
		t.Fatalf("expected S captured, got %v", vars)
	}
}

func TestExtractTypeBoxFromDeclarator_NoInitCall(t *testing.T) {
	d, fi := tbDeclarator(t, "const x = 5;\n")
	vars := map[string]*sitter.Node{}
	extractTypeBoxFromDeclarator(d, fi, vars)
	if len(vars) != 0 {
		t.Fatalf("expected no vars, got %v", vars)
	}
}

func TestExtractTypeBoxFromDeclarator_NotTypeBox(t *testing.T) {
	d, fi := tbDeclarator(t, "const x = Other.Object({});\n")
	vars := map[string]*sitter.Node{}
	extractTypeBoxFromDeclarator(d, fi, vars)
	if len(vars) != 0 {
		t.Fatalf("expected no vars for non-TypeBox, got %v", vars)
	}
}

func TestExtractTypeBoxFromDeclarator_NestedObject(t *testing.T) {
	d, fi := tbDeclarator(t, "const S = Type.Object({ inner: Type.Object({ x: Type.Number() }) });\n")
	vars := map[string]*sitter.Node{}
	extractTypeBoxFromDeclarator(d, fi, vars)
	if vars["S"] == nil {
		t.Fatalf("expected S captured, got %v", vars)
	}
}

func TestExtractTypeBoxFromDeclarator_NoObjectArg(t *testing.T) {
	// Type.Object() with no object literal arg -> objNode nil, nothing captured
	d, fi := tbDeclarator(t, "const x = Type.Object();\n")
	vars := map[string]*sitter.Node{}
	extractTypeBoxFromDeclarator(d, fi, vars)
	if len(vars) != 0 {
		t.Fatalf("expected no vars when no object arg, got %v", vars)
	}
}
