//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveRelativePath_Missing 테스트
package fastify

import "testing"

func TestResolveRelativePath_Missing(t *testing.T) {
	dir := t.TempDir()
	if got := resolveRelativePath(dir, "./missing"); got != "" {
		t.Fatalf("expected empty for missing file, got %q", got)
	}
}
