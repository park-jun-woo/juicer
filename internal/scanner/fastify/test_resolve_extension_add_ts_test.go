//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveExtension_AddTS 테스트
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
