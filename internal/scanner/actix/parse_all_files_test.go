//ff:func feature=scan type=test control=sequence topic=actix
//ff:what parseAllFiles — 유효 파일 파싱과 읽기 실패 스킵을 검증
package actix

import (
	"path/filepath"
	"testing"
)

func TestParseAllFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a.rs", "fn a() {}")
	good := filepath.Join(dir, "a.rs")
	missing := filepath.Join(dir, "missing.rs") // parseFile read error -> skipped

	files := parseAllFiles(dir, []string{good, missing})
	if len(files) != 1 {
		t.Fatalf("expected 1 parsed file, got %d", len(files))
	}
	if filepath.Base(files[0].absPath) != "a.rs" {
		t.Errorf("unexpected parsed file: %s", files[0].absPath)
	}
}
