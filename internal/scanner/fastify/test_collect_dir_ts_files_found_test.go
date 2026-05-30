//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectDirTSFiles_Found 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollectDirTSFiles_Found(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.ts"), []byte(""), 0o644)
	os.WriteFile(filepath.Join(dir, "b.ts"), []byte(""), 0o644)
	os.WriteFile(filepath.Join(dir, "c.js"), []byte(""), 0o644)
	files := collectDirTSFiles(dir)
	if len(files) != 2 {
		t.Fatalf("expected 2 .ts files, got %d: %v", len(files), files)
	}
}
