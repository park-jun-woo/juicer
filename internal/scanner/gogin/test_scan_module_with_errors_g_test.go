//ff:func feature=scan type=test control=sequence
//ff:what TestScan_ModuleWithErrorsG 테스트
package gogin

import "testing"

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
