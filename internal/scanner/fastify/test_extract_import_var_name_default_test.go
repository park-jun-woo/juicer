//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractImportVarName_Default 테스트
package fastify

import "testing"

func TestExtractImportVarName_Default(t *testing.T) {
	fi := mustParse(t, []byte(`import Fastify from "fastify";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	if got := extractImportVarName(stmt, fi.Src); got != "Fastify" {
		t.Fatalf("default import: got %q", got)
	}
}
