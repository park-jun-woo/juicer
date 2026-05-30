//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveRelativePath_NonRelative 테스트
package fastify

import "testing"

func TestResolveRelativePath_NonRelative(t *testing.T) {
	if got := resolveRelativePath("/d", "fastify"); got != "" {
		t.Fatalf("expected empty for non-relative, got %q", got)
	}
}
