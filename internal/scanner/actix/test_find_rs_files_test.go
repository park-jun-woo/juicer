//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestFindRsFiles 테스트
package actix

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFindRsFiles(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main.rs", "fn main() {}")

	writeFile(t, dir, "src/README.md", "doc")

	writeFile(t, dir, "src/util_test.rs", "fn t() {}")

	writeFile(t, dir, "target/build.rs", "fn b() {}")

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
