//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindPyFiles_EmptyDir 테스트
package fastapi

import "testing"

func TestFindPyFiles_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	files, err := findPyFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 0 {
		t.Fatalf("expected 0, got %d", len(files))
	}
}
