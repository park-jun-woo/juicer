//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractNamedImportVarName 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func importClause(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	cs := findAllByType(fi.Root, "import_clause")
	if len(cs) == 0 {
		t.Fatalf("no import_clause in %q", src)
	}
	return cs[0], fi.Src
}

func TestExtractNamedImportVarName_Name(t *testing.T) {
	c, src := importClause(t, `import { join } from "path";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "join" {
		t.Fatalf("got %q, want join", got)
	}
}

func TestExtractNamedImportVarName_Alias(t *testing.T) {
	c, src := importClause(t, `import { join as j } from "path";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "j" {
		t.Fatalf("got %q, want j (alias)", got)
	}
}

func TestExtractNamedImportVarName_NoNamedImports(t *testing.T) {
	// default import clause has no named_imports
	c, src := importClause(t, `import Fastify from "fastify";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "" {
		t.Fatalf("expected empty for default import, got %q", got)
	}
}

func TestExtractNamedImportVarName_MultipleSpecsFirst(t *testing.T) {
	// multiple named imports -> first specifier's name returned
	c, src := importClause(t, `import { readFile, writeFile } from "fs";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "readFile" {
		t.Fatalf("expected readFile (first), got %q", got)
	}
}

func TestExtractNamedImportVarName_EmptyBraces(t *testing.T) {
	// named_imports with no import_specifier -> spec nil -> ""
	c, src := importClause(t, `import {} from "path";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "" {
		t.Fatalf("expected empty for empty braces, got %q", got)
	}
}
