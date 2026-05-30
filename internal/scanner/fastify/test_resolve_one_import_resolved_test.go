//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveOneImport_Resolved 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveOneImport_Resolved(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "mod.ts"), []byte(""), 0o644)
	fi := mustParse(t, []byte(`import m from "./mod";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	imports := map[string]string{}
	resolveOneImport(stmt, fi.Src, dir, imports)
	if imports["m"] == "" {
		t.Fatalf("expected m resolved, got %v", imports)
	}
}
