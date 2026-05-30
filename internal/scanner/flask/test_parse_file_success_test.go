//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseFile_Success 테스트
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_Success(t *testing.T) {
	root := t.TempDir()
	p := filepath.Join(root, "sub", "app.py")
	os.MkdirAll(filepath.Dir(p), 0o755)
	os.WriteFile(p, []byte("x = 1\n"), 0o644)

	fi, err := parseFile(root, p)
	if err != nil {
		t.Fatal(err)
	}
	if fi.relPath != filepath.Join("sub", "app.py") || fi.root == nil {
		t.Fatalf("unexpected fileInfo: %+v", fi)
	}
}
