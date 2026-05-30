//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractAutoloadPrefix_None 테스트
package fastify

import "testing"

func TestExtractAutoloadPrefix_None(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "r" }`)
	if got := extractAutoloadPrefix(obj, src); got != "" {
		t.Fatalf("expected empty prefix, got %q", got)
	}
}
