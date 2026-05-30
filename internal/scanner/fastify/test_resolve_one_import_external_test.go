//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveOneImport_External 테스트
package fastify

import "testing"

func TestResolveOneImport_External(t *testing.T) {

	dir := t.TempDir()
	fi := mustParse(t, []byte(`import Fastify from "fastify";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	imports := map[string]string{}
	resolveOneImport(stmt, fi.Src, dir, imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports for external module, got %v", imports)
	}
}
