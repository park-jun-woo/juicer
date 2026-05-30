//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractImportPath 테스트
package fastify

import "testing"

func TestExtractImportPath(t *testing.T) {
	fi := mustParse(t, []byte(`import Fastify from "fastify";`+"\n"))
	stmts := findAllByType(fi.Root, "import_statement")
	if len(stmts) == 0 {
		t.Fatal("no import_statement")
	}
	if got := extractImportPath(stmts[0], fi.Src); got != "fastify" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractImportPath_NoString(t *testing.T) {
	// a statement without a string child -> ""
	fi := mustParse(t, []byte("const x = 1;\n"))
	stmts := findAllByType(fi.Root, "lexical_declaration")
	if len(stmts) == 0 {
		t.Fatal("no lexical_declaration")
	}
	if got := extractImportPath(stmts[0], fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
