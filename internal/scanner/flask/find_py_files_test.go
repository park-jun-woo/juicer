//ff:func feature=scan type=test control=sequence topic=flask
//ff:what findPyFiles 테스트
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindPyFiles(t *testing.T) {
	root := t.TempDir()
	os.WriteFile(filepath.Join(root, "app.py"), []byte(""), 0o644)
	os.MkdirAll(filepath.Join(root, "pkg"), 0o755)
	os.WriteFile(filepath.Join(root, "pkg", "mod.py"), []byte(""), 0o644)
	os.WriteFile(filepath.Join(root, "readme.txt"), []byte(""), 0o644)
	// excluded dir
	os.MkdirAll(filepath.Join(root, "venv"), 0o755)
	os.WriteFile(filepath.Join(root, "venv", "lib.py"), []byte(""), 0o644)

	files, err := findPyFiles(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 .py files (venv excluded), got %d: %v", len(files), files)
	}
}

func TestFindPyFiles_Error(t *testing.T) {
	if _, err := findPyFiles("/no/such/dir/zzz"); err == nil {
		t.Fatal("expected error for missing root")
	}
}
