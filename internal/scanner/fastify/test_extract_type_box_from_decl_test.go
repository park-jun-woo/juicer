//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractTypeBoxFromDecl 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
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
