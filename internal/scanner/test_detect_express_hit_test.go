//ff:func feature=scan type=test control=sequence
//ff:what TestDetectExpress_Hit 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectExpress_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"),
		[]byte(`{"dependencies":{"express":"^4.18.0"}}`), 0o644)
	if !detectExpress(dir) {
		t.Fatal("expected true")
	}
}
