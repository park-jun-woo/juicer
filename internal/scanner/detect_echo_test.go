//ff:func feature=scan type=test control=sequence
//ff:what TestDetectEcho_Hit go.mod 히트 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectEcho_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("require github.com/labstack/echo/v4 v4.11.0"), 0o644)
	if !detectEcho(dir) {
		t.Fatal("expected true")
	}
}
