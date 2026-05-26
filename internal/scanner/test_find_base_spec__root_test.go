//ff:func feature=scan type=extract control=sequence
//ff:what TestFindBaseSpec_Root 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindBaseSpec_Root(t *testing.T) {
	dir := t.TempDir()
	specPath := filepath.Join(dir, "openapi.yaml")
	os.WriteFile(specPath, []byte("openapi: \"3.0.3\""), 0o644)

	got := FindBaseSpec(dir)
	if got != specPath {
		t.Fatalf("expected %s, got %s", specPath, got)
	}
}
