//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFramework_FastAPIBranch 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFramework_FastAPIBranch(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("fastapi==0.100.0\n"), 0o644)
	if got := DetectFramework(dir); got != "fastapi" {
		t.Fatalf("expected fastapi, got %q", got)
	}
}
