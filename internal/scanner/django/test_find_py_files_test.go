//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestFindPyFiles 테스트
package django

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFindPyFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile2(t, dir, "app/views.py", "x = 1")
	writeFile2(t, dir, "app/README.md", "doc")
	writeFile2(t, dir, "app/views_test.py", "x = 1")
	writeFile2(t, dir, "venv/lib.py", "x = 1")
	writeFile2(t, dir, "app/migrations/0001.py", "x = 1")

	files, err := findPyFiles(dir)
	if err != nil {
		t.Fatalf("findPyFiles error: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("expected exactly 1 file, got %d: %v", len(files), files)
	}
	if filepath.Base(files[0]) != "views.py" {
		t.Errorf("expected views.py, got %s", files[0])
	}
	for _, f := range files {
		if strings.Contains(f, "venv") || strings.Contains(f, "migrations") || strings.Contains(f, "_test") {
			t.Errorf("excluded path leaked: %s", f)
		}
	}
}
