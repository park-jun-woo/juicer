//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestBuildRouterImportMap_SimpleImport 테스트
package fastapi

import (
	"path/filepath"
	"testing"
)

func TestBuildRouterImportMap_SimpleImport(t *testing.T) {
	dir := t.TempDir()
	mkFile(t, dir, "models.py", "x = 1")
	appPath := mkFile(t, dir, "main.py", "x = 1")

	fi := &fileInfo{
		absPath: appPath,
		imports: []importInfo{
			{name: "User", module: ".models"},
			{name: "External", module: "external_pkg"},
		},
	}
	m := buildRouterImportMap(dir, fi, nil)
	if m["User"] != filepath.Join(dir, "models.py") {
		t.Fatalf("User -> %q", m["User"])
	}
	if _, ok := m["External"]; ok {
		t.Fatalf("external import should be unresolved: %v", m)
	}
}
