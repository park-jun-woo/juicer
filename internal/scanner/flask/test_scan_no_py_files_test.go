//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestScan_NoPyFiles 테스트
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_NoPyFiles(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("hi"), 0o644)
	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}
