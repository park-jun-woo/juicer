//ff:func feature=scan type=test control=sequence topic=flask
//ff:what parseAllFiles 테스트
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseAllFiles(t *testing.T) {
	root := t.TempDir()
	p := filepath.Join(root, "app.py")
	os.WriteFile(p, []byte("x = 1\n"), 0o644)

	// one valid file + one missing path (parseFile errors -> continue)
	files := parseAllFiles(root, []string{p, filepath.Join(root, "missing.py")})
	if len(files) != 1 {
		t.Fatalf("expected 1 parsed file, got %d", len(files))
	}
	if files[0].root == nil {
		t.Fatal("expected parsed root node")
	}
}
