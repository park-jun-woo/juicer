//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestParseFileAndAll 테스트
package dotnet

import (
	"path/filepath"
	"testing"
)

func TestParseFileAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UsersController.cs", sampleCtrl)
	fi, err := parseFile(dir, filepath.Join(dir, "UsersController.cs"))
	if err != nil {
		t.Fatal(err)
	}
	if fi.relPath != "UsersController.cs" {
		t.Fatalf("relPath: %q", fi.relPath)
	}
	files := parseAllFiles(dir, []string{
		filepath.Join(dir, "UsersController.cs"),
		filepath.Join(dir, "missing.cs"),
	})
	if len(files) != 1 {
		t.Fatalf("got %d", len(files))
	}
}
