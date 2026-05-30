//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractBasePathFromChain 테스트
package hono

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func valueOfDecl(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	decls := findAllByType(fi.Root, "variable_declarator")
	if len(decls) == 0 {
		t.Fatal("no declarator")
	}
	return decls[0].ChildByFieldName("value"), fi.Src
}

func TestExtractBasePathFromChain_Match(t *testing.T) {
	val, src := valueOfDecl(t, `const a = new Hono().basePath("/api");`+"\n")
	if got := extractBasePathFromChain(val, src); got != "/api" {
		t.Fatalf("got %q, want /api", got)
	}
}

func TestExtractBasePathFromChain_NotCall(t *testing.T) {
	val, src := valueOfDecl(t, `const x = 5;`+"\n")
	if got := extractBasePathFromChain(val, src); got != "" {
		t.Fatalf("expected empty for non-call, got %q", got)
	}
}

func TestExtractBasePathFromChain_NotBasePath(t *testing.T) {
	val, src := valueOfDecl(t, `const a = new Hono().get("/x", h);`+"\n")
	if got := extractBasePathFromChain(val, src); got != "" {
		t.Fatalf("expected empty for non-basePath, got %q", got)
	}
}

func TestExtractBasePathFromChain_NoMemberExpr(t *testing.T) {
	// a plain function call has no member_expression -> ""
	val, src := valueOfDecl(t, `const a = plainCall("/x");`+"\n")
	if got := extractBasePathFromChain(val, src); got != "" {
		t.Fatalf("expected empty for plain call, got %q", got)
	}
}

func TestExtractBasePathFromChain_NoNewHono(t *testing.T) {
	val, src := valueOfDecl(t, `const a = other.basePath("/x");`+"\n")
	if got := extractBasePathFromChain(val, src); got != "" {
		t.Fatalf("expected empty without new Hono, got %q", got)
	}
}
