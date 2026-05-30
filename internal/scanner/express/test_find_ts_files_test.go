//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestFindTSFiles 테스트
package express

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFindTSFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/app.ts", "x")
	writeFile(t, dir, "src/README.md", "doc")
	writeFile(t, dir, "src/app.test.ts", "x")
	writeFile(t, dir, "src/types.d.ts", "x")
	writeFile(t, dir, "node_modules/lib.ts", "x")
	writeFile(t, dir, "tests/it.ts", "x")

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
