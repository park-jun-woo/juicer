//ff:func feature=scan type=test control=sequence topic=django
//ff:what parseAllFiles — 유효 파일 파싱과 읽기 실패 스킵을 검증
package django

import (
	"path/filepath"
	"testing"
)

func TestParseAllFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile2(t, dir, "a.py", "x = 1")
	good := filepath.Join(dir, "a.py")
	missing := filepath.Join(dir, "missing.py") // parseFile error -> skipped

	files := parseAllFiles(dir, []string{good, missing})
	if len(files) != 1 {
		t.Fatalf("expected 1 parsed file, got %d", len(files))
	}
	if filepath.Base(files[0].absPath) != "a.py" {
		t.Errorf("unexpected parsed file: %s", files[0].absPath)
	}
}
