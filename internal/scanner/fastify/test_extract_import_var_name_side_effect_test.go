//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractImportVarName_SideEffect 테스트
package fastify

import "testing"

func TestExtractImportVarName_SideEffect(t *testing.T) {

	fi := mustParse(t, []byte(`import "dotenv/config";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	if got := extractImportVarName(stmt, fi.Src); got != "" {
		t.Fatalf("side-effect import: expected empty, got %q", got)
	}
}
