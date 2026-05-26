//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindTSFiles_NoSrcDir 테스트
package nestjs

import "testing"

func TestFindTSFiles_NoSrcDir(t *testing.T) {
	dir := t.TempDir()
	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 0 {
		t.Fatalf("expected 0, got %d", len(files))
	}
}
