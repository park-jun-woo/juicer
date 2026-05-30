//ff:func feature=scan type=test control=sequence
//ff:what TestDetectActix_Miss 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectActix_Miss(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte("[dependencies]\nrocket = \"0.5\"\n"), 0o644)
	if detectActix(dir) {
		t.Fatal("expected false when actix-web absent")
	}
}
