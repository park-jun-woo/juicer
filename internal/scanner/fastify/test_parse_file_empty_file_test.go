//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestParseFile_EmptyFile 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "empty.ts")
	os.WriteFile(p, []byte(""), 0o644)
	fi, err := parseFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if fi.Root == nil {
		t.Fatal("expected root node for empty file")
	}
}
