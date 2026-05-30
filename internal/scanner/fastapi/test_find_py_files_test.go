//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestFindPyFiles 테스트
package fastapi

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFindPyFiles(t *testing.T) {
	dir := t.TempDir()
	mkFile(t, dir, "app/main.py", "x = 1")
	mkFile(t, dir, "app/README.md", "doc")
	mkFile(t, dir, "app/main_test.py", "x = 1")
	mkFile(t, dir, "venv/lib.py", "x = 1")
	mkFile(t, dir, "tests/it.py", "x = 1")

	files, err := findPyFiles(dir)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if len(files) != 1 || filepath.Base(files[0]) != "main.py" {
		t.Fatalf("expected only main.py, got %v", files)
	}
	for _, f := range files {
		if strings.Contains(f, "venv") || strings.Contains(f, "/tests/") || strings.Contains(f, "_test.") {
			t.Errorf("excluded path leaked: %s", f)
		}
	}
}
