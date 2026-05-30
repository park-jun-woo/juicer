//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractImportVarName_Named 테스트
package fastify

import "testing"

func TestExtractImportVarName_Named(t *testing.T) {
	fi := mustParse(t, []byte(`import { join } from "path";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	if got := extractImportVarName(stmt, fi.Src); got != "join" {
		t.Fatalf("named import: got %q", got)
	}
}
