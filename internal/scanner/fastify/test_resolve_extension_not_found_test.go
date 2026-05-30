//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveExtension_NotFound 테스트
package fastify

import (
	"path/filepath"
	"testing"
)

func TestResolveExtension_NotFound(t *testing.T) {
	dir := t.TempDir()
	if got := resolveExtension(filepath.Join(dir, "missing")); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
