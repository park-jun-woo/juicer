//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindPyFiles_SkipVenv 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindPyFiles_SkipVenv(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "main.py"), []byte("pass"), 0o644)
	venv := filepath.Join(dir, "venv")
	os.MkdirAll(venv, 0o755)
	os.WriteFile(filepath.Join(venv, "activate.py"), []byte("pass"), 0o644)

	files, err := findPyFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("expected 1, got %d", len(files))
	}
}
