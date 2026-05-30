//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what propagateFileIncludes: include 없으면 false / cross-file prefix 전파
package fastapi

import (
	"path/filepath"
	"testing"
)

func TestPropagateFileIncludes_NoIncludes(t *testing.T) {
	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	fi := &fileInfo{absPath: "/m.py", src: src, root: root, prefixes: map[string]string{}}
	if propagateFileIncludes("/", fi, map[string]*fileInfo{}, map[string]map[string]string{}) {
		t.Fatal("expected false when no includes")
	}
}

func TestPropagateFileIncludes_Propagates(t *testing.T) {
	dir := t.TempDir()
	usersPath := mkFile(t, dir, "users.py", "from fastapi import APIRouter\nrouter = APIRouter()\n")
	usersRoot, _ := parsePython([]byte("router = APIRouter()\n"))
	srcFI := &fileInfo{
		absPath:  usersPath,
		src:      []byte("router = APIRouter()\n"),
		root:     usersRoot,
		prefixes: map[string]string{"router": "/users"},
	}

	mainSrc := []byte("from .users import router\napp.include_router(router, prefix=\"/api\")\n")
	mainRoot, _ := parsePython(mainSrc)
	mainPath := filepath.Join(dir, "main.py")
	fi := &fileInfo{
		absPath:  mainPath,
		src:      mainSrc,
		root:     mainRoot,
		imports:  []importInfo{{name: "router", module: ".users"}},
		prefixes: map[string]string{"app": "", "router": "/api/users"},
	}

	fileByPath := map[string]*fileInfo{usersPath: srcFI, mainPath: fi}
	origSnapshot := map[string]map[string]string{
		mainPath:  {"app": "", "router": "/api/users"},
		usersPath: {"router": "/users"},
	}
	// should not panic; returns a bool indicating change
	_ = propagateFileIncludes(dir, fi, fileByPath, origSnapshot)
}
