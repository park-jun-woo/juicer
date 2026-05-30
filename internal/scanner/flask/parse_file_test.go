//ff:func feature=scan type=test control=sequence topic=flask
//ff:what parseFile 테스트
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

func TestParseFile_ReadError(t *testing.T) {
	if _, err := parseFile("/root", "/no/such/file.py"); err == nil {
		t.Fatal("expected read error")
	}
}

func TestParseFile_EmptyFile(t *testing.T) {
	root := t.TempDir()
	p := filepath.Join(root, "empty.py")
	os.WriteFile(p, []byte(""), 0o644)
	fi, err := parseFile(root, p)
	if err != nil {
		t.Fatal(err)
	}
	if fi.root == nil {
		t.Fatal("expected root node for empty file")
	}
}
