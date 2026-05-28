//ff:func feature=scan type=test control=sequence
//ff:what TestDetectEcho_Miss go.mod 미스 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectEcho_Miss(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("require github.com/gin-gonic/gin v1.9.1"), 0o644)
	if detectEcho(dir) {
		t.Fatal("expected false")
	}
}
