//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFramework_FastAPI FastAPI 프레임워크 감지 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFramework_FastAPI(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("fastapi\n"), 0o644)
	if DetectFramework(dir) != "fastapi" {
		t.Fatal("expected fastapi")
	}
}
