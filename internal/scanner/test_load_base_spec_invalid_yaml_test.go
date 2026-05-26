//ff:func feature=scan type=test control=sequence
//ff:what TestLoadBaseSpec_InvalidYAML 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadBaseSpec_InvalidYAML(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "bad.yaml")
	os.WriteFile(f, []byte("{{invalid"), 0o644)
	_, err := LoadBaseSpec(f)
	if err == nil {
		t.Fatal("expected error")
	}
}
