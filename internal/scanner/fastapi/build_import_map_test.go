//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what buildImportMap 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildImportMap(t *testing.T) {
	dir := t.TempDir()
	modelsFile := filepath.Join(dir, "models.py")
	os.WriteFile(modelsFile, []byte("class User: pass"), 0o644)

	imports := []importInfo{
		{name: "User", module: ".models"},
	}
	m := buildImportMap(imports, dir)
	if m["User"] != modelsFile {
		t.Fatalf("expected %s, got %s", modelsFile, m["User"])
	}

	// non-relative import is ignored
	imports2 := []importInfo{{name: "X", module: "app.models"}}
	m2 := buildImportMap(imports2, dir)
	if len(m2) != 0 {
		t.Fatalf("expected 0, got %d", len(m2))
	}
}
