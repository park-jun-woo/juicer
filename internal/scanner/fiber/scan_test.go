//ff:func feature=scan type=test control=sequence
//ff:what Scan — fiber 프로젝트 스캔 진입점 테스트
package fiber

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_PlainModule(t *testing.T) {
	dir := t.TempDir()
	// minimal self-contained module, no external deps
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/m\n\ngo 1.21\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\nfunc main() {}\n"), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	// no fiber routes -> no endpoints, but the full pipeline runs
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}

func TestScan_ModuleWithErrors(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/m\n\ngo 1.21\n"), 0o644)
	// references an undefined symbol -> package loads with type errors,
	// exercising the pkg.Errors warning loop.
	os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\nfunc main() { _ = undefinedSymbol }\n"), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan should tolerate type errors, got: %v", err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}
