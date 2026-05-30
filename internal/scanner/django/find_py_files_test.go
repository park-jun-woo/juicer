//ff:func feature=scan type=test control=sequence topic=django
//ff:what findPyFiles — .py 수집 및 skip 디렉터리/테스트 파일 제외를 검증
package django

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFile2(t *testing.T, dir, rel, content string) {
	t.Helper()
	p := filepath.Join(dir, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestFindPyFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile2(t, dir, "app/views.py", "x = 1")
	writeFile2(t, dir, "app/README.md", "doc")           // non-.py
	writeFile2(t, dir, "app/views_test.py", "x = 1")     // IsTestFile (_test. infix)
	writeFile2(t, dir, "venv/lib.py", "x = 1")           // skipDirs
	writeFile2(t, dir, "app/migrations/0001.py", "x = 1") // skipDirs

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

func TestFindPyFiles_BadRoot(t *testing.T) {
	_, err := findPyFiles(filepath.Join(t.TempDir(), "nope"))
	if err == nil {
		t.Fatal("expected error for missing root")
	}
}
