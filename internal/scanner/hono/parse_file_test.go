//ff:func feature=scan type=test control=sequence topic=hono
//ff:what parseFile 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestParseFile_OK(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "x.ts", "const a = 1\n")
	fi, err := parseFile(filepath.Join(dir, "x.ts"))
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if fi == nil || fi.Root == nil {
		t.Fatal("expected fileInfo with root")
	}
}

func TestParseFile_NotFound(t *testing.T) {
	_, err := parseFile("/no/such/file.ts")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestParseFile_Empty(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "empty.ts", "")
	fi, err := parseFile(filepath.Join(dir, "empty.ts"))
	if err != nil || fi == nil {
		t.Fatalf("empty file should parse: %v", err)
	}
}
