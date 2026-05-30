//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestFindTSFiles 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestFindTSFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/app.ts", "x")
	writeFile(t, dir, "src/types.d.ts", "x")
	writeFile(t, dir, "src/node_modules/lib.ts", "x")
	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel, _ := filepath.Rel(dir, f)
		if rel != filepath.Join("src", "app.ts") {
			t.Errorf("unexpected file: %s", rel)
		}
	}
	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %v", files)
	}
}
