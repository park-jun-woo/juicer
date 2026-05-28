//ff:func feature=scan type=test control=sequence
//ff:what fastify + @nestjs/core 동시 존재 시 감지 거부 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFastify_WithNestJS(t *testing.T) {
	dir := t.TempDir()
	pkg := `{"dependencies":{"fastify":"^4.0.0","@nestjs/core":"^9.0.0"}}`
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(pkg), 0o644)
	if detectFastify(dir) {
		t.Error("expected false when @nestjs/core is also present")
	}
}
