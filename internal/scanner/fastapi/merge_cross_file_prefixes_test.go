//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestMergeCrossFilePrefixes_BoilerplatePattern 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMergeCrossFilePrefixes_BoilerplatePattern(t *testing.T) {
	dir := t.TempDir()

	appDir := filepath.Join(dir, "app")
	sneakersDir := filepath.Join(appDir, "sneakers")
	os.MkdirAll(sneakersDir, 0o755)

	viewsSrc := []byte(`from fastapi import APIRouter, Depends
from app.sneakers.auth import verify_token

router = APIRouter(dependencies=[Depends(verify_token)])

@router.get("/sneakers", status_code=200)
async def list_sneakers():
    pass

@router.get("/sneakers/{sneaker_id}", status_code=200)
async def get_sneaker(sneaker_id: int):
    pass
`)
	viewsPath := filepath.Join(sneakersDir, "views.py")
	os.WriteFile(viewsPath, viewsSrc, 0o644)

	initSrc := []byte(`from fastapi import APIRouter
from app.sneakers.views import router

API_STR = "/api"
sneakers_router = APIRouter(prefix=API_STR)
sneakers_router.include_router(router)
`)
	initPath := filepath.Join(sneakersDir, "__init__.py")
	os.WriteFile(initPath, initSrc, 0o644)

	mainSrc := []byte(`from fastapi import FastAPI
from app.sneakers import sneakers_router

app = FastAPI()
app.include_router(sneakers_router)
`)
	mainPath := filepath.Join(appDir, "main.py")
	os.WriteFile(mainPath, mainSrc, 0o644)

	pyFiles := []string{viewsPath, initPath, mainPath}
	files := parseAllFiles(dir, pyFiles)

	if len(files) != 3 {
		t.Fatalf("expected 3 files, got %d", len(files))
	}

	var viewsFI *fileInfo
	for i := range files {
		if files[i].absPath == viewsPath {
			viewsFI = &files[i]
			break
		}
	}
	if viewsFI == nil {
		t.Fatal("views.py not found in parsed files")
	}

	routerPrefix := viewsFI.prefixes["router"]
	if routerPrefix != "/api" {
		t.Fatalf("expected /api for router in views.py, got %q", routerPrefix)
	}
}
