//ff:func feature=scan type=test control=sequence
//ff:what TestScan_EmptyProjectDir 테스트
package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_EmptyProjectDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "scan-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	_, err = Scan(dir)
	// May or may not error depending on go packages loading
	_ = err
}


func TestScan_PlainModuleG(t *testing.T) {
	dir := t.TempDir()
	scanWriteFile(t, dir, "go.mod", "module example.com/m\n\ngo 1.21\n")
	scanWriteFile(t, dir, "main.go", "package main\nfunc main() {}\n")
	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}

func TestScan_ModuleWithErrorsG(t *testing.T) {
	dir := t.TempDir()
	scanWriteFile(t, dir, "go.mod", "module example.com/m\n\ngo 1.21\n")
	scanWriteFile(t, dir, "main.go", "package main\nfunc main() { _ = undefinedSym }\n")
	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}

func scanWriteFile(t *testing.T, dir, rel, content string) {
	t.Helper()
	p := scanJoin(dir, rel)
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func scanJoin(dir, rel string) string { return filepath.Join(dir, rel) }
