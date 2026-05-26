//ff:func feature=scan type=test control=sequence
//ff:what TestDetectNestJS_NoNest nestjs 미포함 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectNestJS_NoNest(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"dependencies":{"express":"^4.0.0"}}`), 0o644)
	if detectNestJS(dir) {
		t.Fatal("expected false for no nestjs")
	}
}
