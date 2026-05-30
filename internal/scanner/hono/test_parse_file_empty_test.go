//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestParseFile_Empty 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestParseFile_Empty(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "empty.ts", "")
	fi, err := parseFile(filepath.Join(dir, "empty.ts"))
	if err != nil || fi == nil {
		t.Fatalf("empty file should parse: %v", err)
	}
}
