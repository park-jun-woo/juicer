//ff:func feature=scan type=test control=sequence
//ff:what package.json에서 fastify 감지 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFastify(t *testing.T) {
	dir := t.TempDir()
	pkg := `{"dependencies":{"fastify":"^4.0.0"}}`
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(pkg), 0o644)
	if !detectFastify(dir) {
		t.Error("expected true for fastify-only project")
	}
}
