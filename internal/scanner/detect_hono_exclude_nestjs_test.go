//ff:func feature=scan type=test control=sequence
//ff:what detectHono 테스트: @nestjs/core 공존 시 제외
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectHono_ExcludeNestJS(t *testing.T) {
	dir := t.TempDir()
	pkg := `{"dependencies":{"hono":"^4.0.0","@nestjs/core":"^10.0.0"}}`
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(pkg), 0o644)
	if detectHono(dir) {
		t.Error("expected hono not detected when @nestjs/core present")
	}
}
