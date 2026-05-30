//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveRelativePath_NotFound 테스트
package hono

import "testing"

func TestResolveRelativePath_NotFound(t *testing.T) {
	dir := t.TempDir()
	if got := resolveRelativePath(dir, "./missing"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
