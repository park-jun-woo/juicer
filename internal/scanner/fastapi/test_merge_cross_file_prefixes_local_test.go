//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeCrossFilePrefixes_NoImportedRouter 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMergeCrossFilePrefixes_NoImportedRouter(t *testing.T) {
	dir := t.TempDir()

	mainSrc := []byte(`from fastapi import FastAPI, APIRouter

app = FastAPI()
router = APIRouter(prefix="/users")
app.include_router(router)
`)
	mainPath := filepath.Join(dir, "main.py")
	os.WriteFile(mainPath, mainSrc, 0o644)

	pyFiles := []string{mainPath}
	files := parseAllFiles(dir, pyFiles)

	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}

	if files[0].prefixes["router"] != "/users" {
		t.Fatalf("expected /users, got %q", files[0].prefixes["router"])
	}
}
