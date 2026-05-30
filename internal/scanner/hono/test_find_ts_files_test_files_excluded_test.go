//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestFindTSFiles_TestFilesExcluded 테스트
package hono

import "testing"

func TestFindTSFiles_TestFilesExcluded(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", "x")
	writeFile(t, dir, "app.test.ts", "x")
	writeFile(t, dir, "app.spec.ts", "x")

	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel := f[len(dir)+1:]
		if rel != "app.ts" {
			t.Errorf("test file leaked: %s", rel)
		}
	}
}
