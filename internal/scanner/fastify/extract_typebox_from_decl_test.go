//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractTypeBoxFromDecl 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestExtractTypeBoxFromDecl(t *testing.T) {
	src := `const UserSchema = Type.Object({ name: Type.String() });`
	fi := mustParse(t, []byte(src+"\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	if len(decls) == 0 {
		t.Fatal("no lexical_declaration")
	}
	vars := map[string]*sitter.Node{}
	for _, d := range decls {
		extractTypeBoxFromDecl(d, fi, vars)
	}
	if _, ok := vars["UserSchema"]; !ok {
		t.Fatalf("expected UserSchema captured, got %v", vars)
	}
}

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
