//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestBuildRouterImportMap_ChildModule 테스트
package fastapi

import "testing"

func TestBuildRouterImportMap_ChildModule(t *testing.T) {
	dir := t.TempDir()
	mkFile(t, dir, "routers/items.py", "router = APIRouter()")
	appPath := mkFile(t, dir, "main.py", "x = 1")

	fi := &fileInfo{
		absPath: appPath,

		imports: []importInfo{{name: "items", module: ".routers"}},
	}
	includes := []includeCall{
		{parentVar: "app", childVar: "router", childModule: "items"},
		{parentVar: "app", childVar: "x", childModule: ""},
	}
	m := buildRouterImportMap(dir, fi, includes)
	if _, ok := m["items.router"]; !ok {
		t.Fatalf("items.router not resolved: %v", m)
	}

}
