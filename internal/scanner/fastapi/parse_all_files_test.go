//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what parseAllFiles 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseAllFiles(t *testing.T) {
	dir := t.TempDir()
	f1 := filepath.Join(dir, "main.py")
	os.WriteFile(f1, []byte("x = 1\n"), 0o644)
	f2 := filepath.Join(dir, "bad.py")
	// bad file doesn't exist, should be skipped
	files := parseAllFiles(dir, []string{f1, f2})
	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}
}
