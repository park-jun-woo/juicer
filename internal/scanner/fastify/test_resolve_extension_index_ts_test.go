//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveExtension_IndexTS 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveExtension_IndexTS(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "routes")
	os.MkdirAll(sub, 0o755)
	os.WriteFile(filepath.Join(sub, "index.ts"), []byte(""), 0o644)
	if got := resolveExtension(sub); got != filepath.Join(sub, "index.ts") {
		t.Fatalf("got %q", got)
	}
}
