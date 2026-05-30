//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what propagatePrefixPass: 변경 없는 패스 false
package fastapi

import (
	"path/filepath"
	"testing"
)

func TestPropagatePrefixPass_Changes(t *testing.T) {
	dir := t.TempDir()
	usersPath := mkFile(t, dir, "users.py", "from fastapi import APIRouter\nrouter = APIRouter()\n")
	usersRoot, _ := parsePython([]byte("router = APIRouter()\n"))
	srcFI := &fileInfo{
		absPath:  usersPath,
		src:      []byte("router = APIRouter()\n"),
		root:     usersRoot,
		prefixes: map[string]string{"router": ""}, // current prefix empty -> will change
	}

	mainSrc := []byte("from .users import router\napp.include_router(router, prefix=\"/api\")\n")
	mainRoot, _ := parsePython(mainSrc)
	mainPath := filepath.Join(dir, "main.py")
	mainFI := fileInfo{
		absPath:  mainPath,
		src:      mainSrc,
		root:     mainRoot,
		imports:  []importInfo{{name: "router", module: ".users"}},
		prefixes: map[string]string{"app": "", "router": "/api"},
	}

	files := []fileInfo{mainFI}
	fileByPath := map[string]*fileInfo{usersPath: srcFI, mainPath: &files[0]}
	origSnapshot := map[string]map[string]string{
		mainPath:  {"app": "", "router": "/api"},
		usersPath: {"router": ""},
	}
	changed := propagatePrefixPass(dir, files, fileByPath, origSnapshot)
	if !changed {
		t.Fatalf("expected change; srcFI.prefixes=%v", srcFI.prefixes)
	}
}

func TestPropagatePrefixPass_NoChange(t *testing.T) {
	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	files := []fileInfo{
		{absPath: "/a.py", src: src, root: root, prefixes: map[string]string{}},
		{absPath: "/b.py", src: src, root: root, prefixes: map[string]string{}},
	}
	fileByPath := map[string]*fileInfo{}
	if propagatePrefixPass("/", files, fileByPath, map[string]map[string]string{}) {
		t.Fatal("expected false when nothing to propagate")
	}
}
