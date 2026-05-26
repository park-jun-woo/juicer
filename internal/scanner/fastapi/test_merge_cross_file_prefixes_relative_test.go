//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestMergeCrossFilePrefixes_RelativeImport 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMergeCrossFilePrefixes_RelativeImport(t *testing.T) {
	dir := t.TempDir()
	pkgDir := filepath.Join(dir, "pkg")
	os.MkdirAll(pkgDir, 0o755)

	routesSrc := []byte(`from fastapi import APIRouter

router = APIRouter()

@router.get("/items")
async def list_items():
    pass
`)
	routesPath := filepath.Join(pkgDir, "routes.py")
	os.WriteFile(routesPath, routesSrc, 0o644)

	initSrc := []byte(`from fastapi import APIRouter
from .routes import router

main = APIRouter(prefix="/v1")
main.include_router(router)
`)
	initPath := filepath.Join(pkgDir, "__init__.py")
	os.WriteFile(initPath, initSrc, 0o644)

	pyFiles := []string{routesPath, initPath}
	files := parseAllFiles(dir, pyFiles)

	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}

	var routesFI *fileInfo
	for i := range files {
		if files[i].absPath == routesPath {
			routesFI = &files[i]
			break
		}
	}
	if routesFI == nil {
		t.Fatal("routes.py not found")
	}

	routerPrefix := routesFI.prefixes["router"]
	if routerPrefix != "/v1" {
		t.Fatalf("expected /v1 for router in routes.py, got %q", routerPrefix)
	}
}
