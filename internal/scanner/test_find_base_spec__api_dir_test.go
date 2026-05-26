//ff:func feature=scan type=extract control=sequence
//ff:what TestFindBaseSpec_ApiDir 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindBaseSpec_ApiDir(t *testing.T) {
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "api")
	os.MkdirAll(apiDir, 0o755)
	specPath := filepath.Join(apiDir, "openapi.yaml")
	os.WriteFile(specPath, []byte("openapi: \"3.0.3\""), 0o644)

	got := FindBaseSpec(dir)
	if got != specPath {
		t.Fatalf("expected %s, got %s", specPath, got)
	}
}
