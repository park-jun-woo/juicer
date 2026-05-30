//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveOneImport_Relative 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveOneImport_Relative(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "users.ts", "x")
	fi := mustParse(t, []byte(`import r from './users';`))
	imports := map[string]string{}
	resolveOneImport(firstImportStmt(t, fi), fi.Src, dir, imports, dir, nil)
	if imports["r"] != filepath.Join(dir, "users.ts") {
		t.Fatalf("got %v", imports)
	}
}
