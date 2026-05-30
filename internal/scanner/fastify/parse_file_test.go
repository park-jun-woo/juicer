//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what parseFile 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_Success(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "app.ts")
	os.WriteFile(p, []byte("const app = Fastify();\n"), 0o644)
	fi, err := parseFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if fi.Path != p || fi.Root == nil {
		t.Fatalf("unexpected fileInfo: %+v", fi)
	}
}

func TestParseFile_ReadError(t *testing.T) {
	if _, err := parseFile("/no/such/file.ts"); err == nil {
		t.Fatal("expected read error")
	}
}

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
