//ff:func feature=scan type=test control=sequence
//ff:what TestDetectExpress_ExcludedByNest 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectExpress_ExcludedByNest(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"),
		[]byte(`{"dependencies":{"express":"^4","@nestjs/core":"^10"}}`), 0o644)
	if detectExpress(dir) {
		t.Fatal("expected false when @nestjs/core present")
	}
}
