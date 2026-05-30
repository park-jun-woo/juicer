//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestParseFileAndAll 테스트
package spring

import (
	"path/filepath"
	"testing"
)

func TestParseFileAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserController.java", sampleController)
	fi, err := parseFile(dir, filepath.Join(dir, "UserController.java"))
	if err != nil {
		t.Fatal(err)
	}
	if fi.relPath != "UserController.java" {
		t.Fatalf("relPath: %q", fi.relPath)
	}
	files := parseAllFiles(dir, []string{
		filepath.Join(dir, "UserController.java"),
		filepath.Join(dir, "missing.java"),
	})
	if len(files) != 1 {
		t.Fatalf("got %d", len(files))
	}
}
