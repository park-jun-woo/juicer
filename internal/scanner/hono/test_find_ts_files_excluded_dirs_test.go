//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestFindTSFiles_ExcludedDirs 테스트
package hono

import "testing"

func TestFindTSFiles_ExcludedDirs(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", "x")
	writeFile(t, dir, "node_modules/lib.ts", "x")
	writeFile(t, dir, "dist/out.ts", "x")
	writeFile(t, dir, "build/out.ts", "x")
	writeFile(t, dir, ".git/hook.ts", "x")
	writeFile(t, dir, "__tests__/case.ts", "x")

	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel := f[len(dir)+1:]
		if rel != "app.ts" {
			t.Errorf("unexpected collected file: %s", rel)
		}
	}
	if len(files) != 1 {
		t.Fatalf("expected only app.ts, got %v", files)
	}
}
