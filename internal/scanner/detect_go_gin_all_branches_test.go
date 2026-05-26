//ff:func feature=scan type=test control=sequence
//ff:what TestDetectGoGin_AllBranches miss/no-gin 전 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectGoGin_AllBranches(t *testing.T) {
	// no go.mod file
	dir1 := t.TempDir()
	if detectGoGin(dir1) {
		t.Fatal("expected false when no go.mod")
	}

	// go.mod without gin
	dir2 := t.TempDir()
	os.WriteFile(filepath.Join(dir2, "go.mod"), []byte("module example.com/test\ngo 1.21\n"), 0o644)
	if detectGoGin(dir2) {
		t.Fatal("expected false when gin not in go.mod")
	}
}
