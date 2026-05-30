//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestScan_NonEchoModule_Round5 테스트
package echo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_NonEchoModule_Round5(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module x\ngo 1.21\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "m.go"), []byte("package x\nfunc Plain() {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	res, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Endpoints) != 0 {
		t.Fatalf("expected no echo endpoints, got %d", len(res.Endpoints))
	}
}
