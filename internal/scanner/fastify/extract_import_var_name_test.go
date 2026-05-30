//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractImportVarName 테스트
package fastify

import "testing"

func TestExtractImportVarName_Default(t *testing.T) {
	fi := mustParse(t, []byte(`import Fastify from "fastify";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	if got := extractImportVarName(stmt, fi.Src); got != "Fastify" {
		t.Fatalf("default import: got %q", got)
	}
}

func TestExtractImportVarName_Named(t *testing.T) {
	fi := mustParse(t, []byte(`import { join } from "path";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	if got := extractImportVarName(stmt, fi.Src); got != "join" {
		t.Fatalf("named import: got %q", got)
	}
}

func TestExtractImportVarName_SideEffect(t *testing.T) {
	// side-effect import has no import_clause -> ""
	fi := mustParse(t, []byte(`import "dotenv/config";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	if got := extractImportVarName(stmt, fi.Src); got != "" {
		t.Fatalf("side-effect import: expected empty, got %q", got)
	}
}
