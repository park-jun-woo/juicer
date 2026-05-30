//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagatePrefixToRouteFiles 테스트
package fastapi

import (
	"path/filepath"
	"testing"
)

func TestPropagatePrefixToRouteFiles(t *testing.T) {
	dir := t.TempDir()
	usersPath := mkFile(t, dir, "users.py", "from fastapi import APIRouter\nrouter = APIRouter()\n")
	mainPath := filepath.Join(dir, "main.py")
	mkFile(t, dir, "main.py", "from .users import router\napp.include_router(router, prefix=\"/api\")\n")

	usersRoot, _ := parsePython([]byte("router = APIRouter()\n"))
	mainSrc := []byte("from .users import router\napp.include_router(router, prefix=\"/api\")\n")
	mainRoot, _ := parsePython(mainSrc)

	files := []fileInfo{
		{
			absPath:  usersPath,
			src:      []byte("router = APIRouter()\n"),
			root:     usersRoot,
			prefixes: map[string]string{"router": ""},
		},
		{
			absPath:  mainPath,
			src:      mainSrc,
			root:     mainRoot,
			imports:  []importInfo{{name: "router", module: ".users"}},
			prefixes: map[string]string{"app": "", "router": "/api"},
		},
	}

	propagatePrefixToRouteFiles(dir, files)

	if files[0].prefixes["router"] == "" {
		t.Logf("router prefix after propagation: %q", files[0].prefixes["router"])
	}
}
