//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseFile_OK 테스트
package actix

import (
	"path/filepath"
	"testing"
)

func TestParseFile_OK(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.rs", "fn main() {}")
	abs := filepath.Join(dir, "src/main.rs")

	fi, err := parseFile(dir, abs)
	if err != nil {
		t.Fatalf("parseFile error: %v", err)
	}
	if fi.relPath != filepath.Join("src", "main.rs") {
		t.Errorf("relPath = %q", fi.relPath)
	}
	if fi.root == nil {
		t.Error("expected non-nil root")
	}
	if fi.projectRoot != dir {
		t.Errorf("projectRoot = %q, want %q", fi.projectRoot, dir)
	}
}
