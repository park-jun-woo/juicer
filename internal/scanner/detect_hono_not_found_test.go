//ff:func feature=scan type=test control=sequence
//ff:what detectHono 테스트: hono 미설치
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectHono_NotFound(t *testing.T) {
	dir := t.TempDir()
	pkg := `{"dependencies":{"express":"^4.0.0"}}`
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(pkg), 0o644)
	if detectHono(dir) {
		t.Error("expected hono not detected")
	}
}
