//ff:func feature=scan type=test control=sequence
//ff:what TestDetectNestJS_Hit package.json 히트 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectNestJS_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"dependencies":{"@nestjs/core":"^10"}}`), 0o644)
	if !detectNestJS(dir) {
		t.Fatal("expected true")
	}
}
