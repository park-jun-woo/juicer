//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestParseFile_OK 테스트
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
