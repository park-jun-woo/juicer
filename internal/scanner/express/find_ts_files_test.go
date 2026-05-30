//ff:func feature=scan type=test control=sequence topic=express
//ff:what findTSFiles: .ts 수집 + node_modules/test디렉터리/테스트파일/비소스 제외 + 잘못된 root 에러
package express

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFindTSFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/app.ts", "x")
	writeFile(t, dir, "src/README.md", "doc")    // non-source
	writeFile(t, dir, "src/app.test.ts", "x")    // test file excluded
	writeFile(t, dir, "src/types.d.ts", "x")     // .d.ts excluded
	writeFile(t, dir, "node_modules/lib.ts", "x") // skip dir
	writeFile(t, dir, "tests/it.ts", "x")        // test dir excluded

	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if len(files) != 1 || filepath.Base(files[0]) != "app.ts" {
		t.Fatalf("expected only app.ts, got %v", files)
	}
	for _, f := range files {
		if strings.Contains(f, "node_modules") || strings.Contains(f, "/tests/") || strings.Contains(f, ".test.") {
			t.Errorf("excluded path leaked: %s", f)
		}
	}
}

func TestFindTSFiles_BadRoot(t *testing.T) {
	_, err := findTSFiles(filepath.Join(t.TempDir(), "nope"))
	if err == nil {
		t.Fatal("expected error for non-existent root")
	}
}
