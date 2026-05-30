//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractAutoloadDir_EmptySegments 테스트
package fastify

import "testing"

func TestExtractAutoloadDir_EmptySegments(t *testing.T) {

	obj, src := firstObject(t, `{ dir: resolve(base) }`)
	if got := extractAutoloadDir(obj, src); got != "" {
		t.Fatalf("expected empty for no string segments, got %q", got)
	}
}
