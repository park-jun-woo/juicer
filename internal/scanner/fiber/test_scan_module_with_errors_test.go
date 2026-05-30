//ff:func feature=scan type=test control=sequence
//ff:what TestScan_ModuleWithErrors 테스트
package fiber

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_ModuleWithErrors(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/m\n\ngo 1.21\n"), 0o644)

	os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\nfunc main() { _ = undefinedSymbol }\n"), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan should tolerate type errors, got: %v", err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}
