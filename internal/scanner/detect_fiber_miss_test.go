//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFiber_Miss go.mod에 fiber가 없는 경우 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFiber_Miss(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("require github.com/gin-gonic/gin v1.9.1"), 0o644)
	if detectFiber(dir) {
		t.Fatal("expected false")
	}
}
