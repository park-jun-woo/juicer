//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveOneImport_Alias 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveOneImport_Alias(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/users.ts", "x")
	fi := mustParse(t, []byte(`import r from '@app/users';`))
	imports := map[string]string{}
	aliases := map[string]string{"@app/": filepath.Join("src") + string(filepath.Separator)}
	resolveOneImport(firstImportStmt(t, fi), fi.Src, dir, imports, dir, aliases)
	if imports["r"] == "" {
		t.Fatalf("expected alias resolution, got %v", imports)
	}
}
