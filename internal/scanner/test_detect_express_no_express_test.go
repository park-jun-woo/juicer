//ff:func feature=scan type=test control=sequence
//ff:what TestDetectExpress_NoExpress 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectExpress_NoExpress(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"dependencies":{}}`), 0o644)
	if detectExpress(dir) {
		t.Fatal("expected false without express")
	}
}
