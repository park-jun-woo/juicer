//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveImportPath 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveImportPath(t *testing.T) {
	dir := t.TempDir()
	modelsFile := filepath.Join(dir, "models.py")
	os.WriteFile(modelsFile, []byte("pass"), 0o644)

	got := resolveImportPath(dir, ".models")
	if got != modelsFile {
		t.Fatalf("expected %s, got %s", modelsFile, got)
	}

	// non-relative
	got2 := resolveImportPath(dir, "app.models")
	if got2 != "" {
		t.Fatalf("expected empty for non-relative, got %s", got2)
	}

	// dots only (no module part)
	got3 := resolveImportPath(dir, ".")
	if got3 != "" {
		t.Fatalf("expected empty for dots only, got %s", got3)
	}

	// __init__.py fallback
	pkgDir := filepath.Join(dir, "schemas")
	os.MkdirAll(pkgDir, 0o755)
	os.WriteFile(filepath.Join(pkgDir, "__init__.py"), []byte("pass"), 0o644)
	got4 := resolveImportPath(dir, ".schemas")
	if got4 != filepath.Join(pkgDir, "__init__.py") {
		t.Fatalf("expected __init__.py, got %s", got4)
	}
}
