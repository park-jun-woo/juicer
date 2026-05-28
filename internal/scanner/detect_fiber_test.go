//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFiber_Hit go.mod 히트 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFiber_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("require github.com/gofiber/fiber/v2 v2.52.0"), 0o644)
	if !detectFiber(dir) {
		t.Fatal("expected true")
	}
}
