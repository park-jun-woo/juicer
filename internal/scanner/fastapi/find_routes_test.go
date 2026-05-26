//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindPyFiles_Basic 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindPyFiles_Basic(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "main.py"), []byte("pass"), 0o644)
	os.WriteFile(filepath.Join(dir, "routes.py"), []byte("pass"), 0o644)
	os.WriteFile(filepath.Join(dir, "readme.md"), []byte("# hi"), 0o644)

	files, err := findPyFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2, got %d", len(files))
	}
}
