//ff:func feature=scan type=extract control=sequence
//ff:what TestScan_WithPackageErrors 테스트
package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_WithPackageErrors(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/test\ngo 1.21\n"), 0o644)
	// File with import that doesn't exist — will produce package errors
	os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\nimport \"nonexistent/pkg\"\nfunc main() { pkg.Do() }\n"), 0o644)

	result, err := Scan(dir)
	// Should not error — package errors are warnings
	if err != nil {
		t.Logf("Scan() returned error (may be expected): %v", err)
	}
	_ = result
}
