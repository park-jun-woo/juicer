//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFramework_NestJSBranch 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFramework_NestJSBranch(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"dependencies":{"@nestjs/core":"^10"}}`), 0o644)
	if got := DetectFramework(dir); got != "nestjs" {
		t.Fatalf("expected nestjs, got %q", got)
	}
}
