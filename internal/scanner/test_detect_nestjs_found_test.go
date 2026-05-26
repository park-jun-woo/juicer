//ff:func feature=scan type=test control=sequence
//ff:what TestDetectNestJS_Found package.json 감지 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectNestJS_Found(t *testing.T) {
	dir := t.TempDir()
	pkg := `{"dependencies":{"@nestjs/core":"^10.0.0"}}`
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(pkg), 0o644)
	if !detectNestJS(dir) {
		t.Fatal("expected true")
	}
}
