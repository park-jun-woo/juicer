//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveExtension_AlreadyTS 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveExtension_AlreadyTS(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "app.ts")
	os.WriteFile(p, []byte(""), 0o644)
	if got := resolveExtension(p); got != p {
		t.Fatalf("got %q, want %q", got, p)
	}
}
