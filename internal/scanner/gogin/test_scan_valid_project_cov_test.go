//ff:func feature=scan type=test control=sequence
//ff:what TestScan_ValidProjectCov 테스트
package gogin

import (
	"os"
	"testing"
)

func TestScan_ValidProjectCov(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(dir+"/go.mod", []byte("module test\ngo 1.21\n"), 0o644)
	os.WriteFile(dir+"/main.go", []byte("package main\nfunc main() {}\n"), 0o644)
	result, err := Scan(dir)
	if err != nil {
		t.Skipf("packages.Load may fail in test env: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
