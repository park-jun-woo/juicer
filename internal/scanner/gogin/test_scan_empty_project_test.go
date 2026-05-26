//ff:func feature=scan type=extract control=sequence
//ff:what TestScan_EmptyProject 테스트
package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_EmptyProject(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/test\ngo 1.21\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\nfunc main() {}\n"), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
