//ff:func feature=scan type=extract control=sequence
//ff:what TestFindBaseSpec_DocsDir 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindBaseSpec_DocsDir(t *testing.T) {
	dir := t.TempDir()
	docsDir := filepath.Join(dir, "docs")
	os.MkdirAll(docsDir, 0o755)
	specPath := filepath.Join(docsDir, "openapi.yml")
	os.WriteFile(specPath, []byte("openapi: \"3.0.3\""), 0o644)

	got := FindBaseSpec(dir)
	if got != specPath {
		t.Fatalf("expected %s, got %s", specPath, got)
	}
}
