//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what resolveRelativePath 테스트
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

func TestResolveRelativePath_NonRelative(t *testing.T) {
	if got := resolveRelativePath("/d", "fastify"); got != "" {
		t.Fatalf("expected empty for non-relative, got %q", got)
	}
}

func TestResolveRelativePath_Missing(t *testing.T) {
	dir := t.TempDir()
	if got := resolveRelativePath(dir, "./missing"); got != "" {
		t.Fatalf("expected empty for missing file, got %q", got)
	}
}
