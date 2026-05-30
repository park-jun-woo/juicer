//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveRelativePath_Resolved 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveRelativePath_Resolved(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "mod.ts"), []byte(""), 0o644)
	got := resolveRelativePath(dir, "./mod")
	if got != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %q", got)
	}
}
