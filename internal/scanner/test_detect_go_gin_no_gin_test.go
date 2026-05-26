//ff:func feature=scan type=test control=sequence
//ff:what TestDetectGoGin_NoGin gin 미포함 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectGoGin_NoGin(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module test\n"), 0o644)
	if detectGoGin(dir) {
		t.Fatal("expected false for no gin")
	}
}
