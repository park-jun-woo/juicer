//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what buildRouterImportMap: 단순 import 매핑 + childModule include 처리 분기
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func mkFile(t *testing.T, dir, rel, content string) string {
	t.Helper()
	p := filepath.Join(dir, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestBuildRouterImportMap_SimpleImport(t *testing.T) {
	dir := t.TempDir()
	mkFile(t, dir, "models.py", "x = 1")
	appPath := mkFile(t, dir, "main.py", "x = 1")

	fi := &fileInfo{
		absPath: appPath,
		imports: []importInfo{
			{name: "User", module: ".models"},
			{name: "External", module: "external_pkg"}, // unresolved
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

func TestBuildRouterImportMap_ChildModule(t *testing.T) {
	dir := t.TempDir()
	mkFile(t, dir, "routers/items.py", "router = APIRouter()")
	appPath := mkFile(t, dir, "main.py", "x = 1")

	fi := &fileInfo{
		absPath: appPath,
		// from .routers import items  => name "items", module ".routers"
		imports: []importInfo{{name: "items", module: ".routers"}},
	}
	includes := []includeCall{
		{parentVar: "app", childVar: "router", childModule: "items"},
		{parentVar: "app", childVar: "x", childModule: ""}, // empty childModule -> skipped
	}
	m := buildRouterImportMap(dir, fi, includes)
	if _, ok := m["items.router"]; !ok {
		t.Fatalf("items.router not resolved: %v", m)
	}
	// The "qualifiedKey already exists" continue branch is unreachable in practice:
	// import names never contain a dot, so they cannot collide with childModule.childVar.
}
