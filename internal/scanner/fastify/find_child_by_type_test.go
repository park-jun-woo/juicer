//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what findChildByType 테스트
package fastify

import "testing"

func TestFindChildByType(t *testing.T) {
	fi := mustParse(t, []byte("const x = 1;\n"))
	// program -> lexical_declaration exists as direct child
	decl := findChildByType(fi.Root, "lexical_declaration")
	if decl == nil {
		t.Fatal("expected lexical_declaration child")
	}
	// type that is not a direct child -> nil
	if got := findChildByType(fi.Root, "object"); got != nil {
		t.Fatalf("expected nil for absent type, got %s", got.Type())
	}
}
