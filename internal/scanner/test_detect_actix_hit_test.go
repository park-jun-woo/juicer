//ff:func feature=scan type=test control=sequence
//ff:what TestDetectActix_Hit 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectActix_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte("[dependencies]\nactix-web = \"4\"\n"), 0o644)
	if !detectActix(dir) {
		t.Fatal("expected true")
	}
}
