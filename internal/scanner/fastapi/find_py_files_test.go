//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findPyFiles: .py 수집 + venv/test디렉터리/테스트파일/비py 제외 + 잘못된 root 에러
package fastapi

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFindPyFiles(t *testing.T) {
	dir := t.TempDir()
	mkFile(t, dir, "app/main.py", "x = 1")
	mkFile(t, dir, "app/README.md", "doc")        // non-py
	mkFile(t, dir, "app/main_test.py", "x = 1")    // test file excluded (_test.)
	mkFile(t, dir, "venv/lib.py", "x = 1")         // skip dir
	mkFile(t, dir, "tests/it.py", "x = 1")         // test dir excluded

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

func TestFindPyFiles_BadRoot(t *testing.T) {
	_, err := findPyFiles(filepath.Join(t.TempDir(), "nope"))
	if err == nil {
		t.Fatal("expected error for non-existent root")
	}
}
