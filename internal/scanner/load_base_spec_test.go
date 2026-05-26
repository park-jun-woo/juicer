//ff:func feature=scan type=test control=sequence
//ff:what TestLoadBaseSpec_ValidCov 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadBaseSpec_ValidCov(t *testing.T) {
	dir := t.TempDir()
	spec := "openapi: \"3.0.0\"\ninfo:\n  title: test\n  version: \"1.0\"\npaths: {}\n"
	f := filepath.Join(dir, "openapi.yaml")
	os.WriteFile(f, []byte(spec), 0o644)
	node, err := LoadBaseSpec(f)
	if err != nil {
		t.Fatal(err)
	}
	if node == nil {
		t.Fatal("expected non-nil node")
	}
}
