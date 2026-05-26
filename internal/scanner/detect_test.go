//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFramework_GoGinBranch GoGin 감지 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFramework_GoGinBranch(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("require github.com/gin-gonic/gin v1.9.1"), 0o644)
	if got := DetectFramework(dir); got != "gogin" {
		t.Fatalf("expected gogin, got %q", got)
	}
}
