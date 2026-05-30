//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeImportedRouterPrefixes_Merges 테스트
package fastapi

import (
	"path/filepath"
	"testing"
)

func TestMergeImportedRouterPrefixes_Merges(t *testing.T) {
	dir := t.TempDir()
	usersPath := mkFile(t, dir, "users.py", "from fastapi import APIRouter\nrouter = APIRouter()\n")
	mainSrc := []byte("from .users import router as users_router\napp.include_router(users_router, prefix=\"/api\")\n")
	root, err := parsePython(mainSrc)
	if err != nil {
		t.Fatal(err)
	}
	fi := fileInfo{
		absPath: filepath.Join(dir, "main.py"),
		src:     mainSrc,
		root:    root,
		imports: []importInfo{{name: "users_router", module: ".users"}},
		prefixes: map[string]string{
			"app": "",
		},
	}
	globalPrefixes := map[string]map[string]string{
		usersPath: {"users_router": "/users"},
	}
	files := []fileInfo{fi}
	mergeImportedRouterPrefixes(dir, files, globalPrefixes)

	if _, ok := files[0].prefixes["users_router"]; !ok {
		t.Fatalf("expected users_router prefix merged, got %v", files[0].prefixes)
	}
}
