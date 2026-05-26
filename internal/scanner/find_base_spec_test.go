//ff:func feature=scan type=test control=sequence
//ff:what TestFindBaseSpec_Found 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindBaseSpec_Found(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "openapi.yaml"), []byte("openapi: 3.0.0"), 0o644)
	result := FindBaseSpec(dir)
	if result == "" {
		t.Fatal("expected to find openapi.yaml")
	}

	// not found
	dir2 := t.TempDir()
	if FindBaseSpec(dir2) != "" {
		t.Fatal("expected empty for missing spec")
	}
}

