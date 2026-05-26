//ff:func feature=scan type=test control=sequence
//ff:what TestDetectGoGin_Hit go.mod 히트 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectGoGin_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("require github.com/gin-gonic/gin v1.9.1"), 0o644)
	if !detectGoGin(dir) {
		t.Fatal("expected true")
	}
}
