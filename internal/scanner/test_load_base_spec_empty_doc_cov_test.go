//ff:func feature=scan type=test control=sequence
//ff:what TestLoadBaseSpec_EmptyDocCov 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadBaseSpec_EmptyDocCov(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "empty.yaml")
	// Empty file produces empty document
	os.WriteFile(f, []byte(""), 0o644)
	_, err := LoadBaseSpec(f)
	if err == nil {
		t.Fatal("expected error for empty document")
	}
}
