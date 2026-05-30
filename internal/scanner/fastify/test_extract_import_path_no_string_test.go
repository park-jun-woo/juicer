//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractImportPath_NoString 테스트
package fastify

import "testing"

func TestExtractImportPath_NoString(t *testing.T) {

	fi := mustParse(t, []byte("const x = 1;\n"))
	stmts := findAllByType(fi.Root, "lexical_declaration")
	if len(stmts) == 0 {
		t.Fatal("no lexical_declaration")
	}
	if got := extractImportPath(stmts[0], fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
