//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findRsFiles — .rs 수집, skip 디렉터리/테스트 디렉터리/테스트 파일 제외를 검증
package actix

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFindRsFiles(t *testing.T) {
	dir := t.TempDir()
	// Included .rs file.
	writeFile(t, dir, "src/main.rs", "fn main() {}")
	// Non-.rs file is ignored.
	writeFile(t, dir, "src/README.md", "doc")
	// Test file (IsTestFile) is excluded.
	writeFile(t, dir, "src/util_test.rs", "fn t() {}")
	// File inside a skipDirs directory ("target") is excluded.
	writeFile(t, dir, "target/build.rs", "fn b() {}")
	// File inside a test directory ("tests") is excluded.
	writeFile(t, dir, "tests/it.rs", "fn it() {}")

	files, err := findRsFiles(dir)
	if err != nil {
		t.Fatalf("findRsFiles error: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("expected exactly 1 file, got %d: %v", len(files), files)
	}
	if filepath.Base(files[0]) != "main.rs" {
		t.Errorf("expected main.rs, got %s", files[0])
	}
	for _, f := range files {
		if strings.Contains(f, "target") || strings.Contains(f, "/tests/") || strings.Contains(f, "_test") {
			t.Errorf("excluded path leaked: %s", f)
		}
	}
}

func TestFindRsFiles_BadRoot(t *testing.T) {
	// A non-existent root makes filepath.Walk pass an error into the callback.
	_, err := findRsFiles(filepath.Join(t.TempDir(), "does-not-exist"))
	if err == nil {
		t.Fatal("expected error for non-existent root")
	}
}
