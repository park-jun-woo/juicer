//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFlask_Hit 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFlask_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("Flask==3.0\n"), 0o644)
	if !detectFlask(dir) {
		t.Fatal("expected true")
	}
}
