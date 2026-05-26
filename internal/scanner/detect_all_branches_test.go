//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFramework_AllBranches nestjs/fastapi/unknown 전 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFramework_AllBranches(t *testing.T) {
	// nestjs
	dir1 := t.TempDir()
	os.WriteFile(filepath.Join(dir1, "package.json"), []byte(`{"dependencies":{"@nestjs/core":"*"}}`), 0o644)
	if got := DetectFramework(dir1); got != "nestjs" {
		t.Fatalf("expected nestjs, got %q", got)
	}

	// fastapi
	dir2 := t.TempDir()
	os.WriteFile(filepath.Join(dir2, "requirements.txt"), []byte("fastapi\n"), 0o644)
	if got := DetectFramework(dir2); got != "fastapi" {
		t.Fatalf("expected fastapi, got %q", got)
	}

	// unknown
	dir3 := t.TempDir()
	if got := DetectFramework(dir3); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
