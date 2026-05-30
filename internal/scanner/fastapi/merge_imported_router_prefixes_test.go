//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what mergeImportedRouterPrefixes: import 라우터 prefix 병합 / include 없으면 스킵
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

	// users_router should now have a merged prefix in files[0]
	if _, ok := files[0].prefixes["users_router"]; !ok {
		t.Fatalf("expected users_router prefix merged, got %v", files[0].prefixes)
	}
}

func TestMergeImportedRouterPrefixes_NoIncludes(t *testing.T) {
	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	fi := fileInfo{absPath: "/m.py", src: src, root: root, prefixes: map[string]string{}}
	files := []fileInfo{fi}
	// should not panic and leave prefixes untouched
	mergeImportedRouterPrefixes("/", files, map[string]map[string]string{})
	if len(files[0].prefixes) != 0 {
		t.Fatalf("expected no change, got %v", files[0].prefixes)
	}
}
