//ff:func feature=scan type=extract control=sequence
//ff:what TestScan_NonGoDir 테스트
package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_NonGoDir(t *testing.T) {
	// A directory with no go files — packages.Load may return empty
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("not go"), 0o644)
	result, err := Scan(dir)
	// packages.Load on non-Go dir may error
	if err != nil {
		// Expected — packages.Load can't find Go files
		return
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
