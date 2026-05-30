//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestParseFileAndAll 테스트
package quarkus

import (
	"path/filepath"
	"testing"
)

func TestParseFileAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserResource.java", sampleResource)
	fi, err := parseFile(dir, filepath.Join(dir, "UserResource.java"))
	if err != nil {
		t.Fatal(err)
	}
	if fi.relPath != "UserResource.java" {
		t.Fatalf("relPath: %q", fi.relPath)
	}
	files := parseAllFiles(dir, []string{
		filepath.Join(dir, "UserResource.java"),
		filepath.Join(dir, "missing.java"),
	})
	if len(files) != 1 {
		t.Fatalf("expected 1 parsed, got %d", len(files))
	}
}
