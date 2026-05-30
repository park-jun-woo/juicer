//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what resolveExtension 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveExtension_AddTS(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "users.ts"), []byte(""), 0o644)
	base := filepath.Join(dir, "users")
	if got := resolveExtension(base); got != base+".ts" {
		t.Fatalf("got %q, want %q", got, base+".ts")
	}
}

func TestResolveExtension_IndexTS(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "routes")
	os.MkdirAll(sub, 0o755)
	os.WriteFile(filepath.Join(sub, "index.ts"), []byte(""), 0o644)
	if got := resolveExtension(sub); got != filepath.Join(sub, "index.ts") {
		t.Fatalf("got %q", got)
	}
}

func TestResolveExtension_AlreadyTS(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "app.ts")
	os.WriteFile(p, []byte(""), 0o644)
	if got := resolveExtension(p); got != p {
		t.Fatalf("got %q, want %q", got, p)
	}
}

func TestResolveExtension_NotFound(t *testing.T) {
	dir := t.TempDir()
	if got := resolveExtension(filepath.Join(dir, "missing")); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
