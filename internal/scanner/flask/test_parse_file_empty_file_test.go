//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseFile_EmptyFile 테스트
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

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
