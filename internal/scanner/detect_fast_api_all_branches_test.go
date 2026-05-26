//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFastAPI_AllBranches pyproject/empty/miss 전 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFastAPI_AllBranches(t *testing.T) {
	// pyproject.toml hit
	dir1 := t.TempDir()
	os.WriteFile(filepath.Join(dir1, "pyproject.toml"), []byte("[tool.poetry.dependencies]\nfastapi = \"*\"\n"), 0o644)
	if !detectFastAPI(dir1) {
		t.Fatal("expected true for fastapi in pyproject.toml")
	}

	// no files at all
	dir2 := t.TempDir()
	if detectFastAPI(dir2) {
		t.Fatal("expected false when no files present")
	}

	// file exists but no fastapi
	dir3 := t.TempDir()
	os.WriteFile(filepath.Join(dir3, "requirements.txt"), []byte("flask==2.0.0\n"), 0o644)
	if detectFastAPI(dir3) {
		t.Fatal("expected false when fastapi not in file")
	}
}
