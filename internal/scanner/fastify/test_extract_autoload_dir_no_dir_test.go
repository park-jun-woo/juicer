//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractAutoloadDir_NoDir 테스트
package fastify

import "testing"

func TestExtractAutoloadDir_NoDir(t *testing.T) {
	obj, src := firstObject(t, `{ options: { prefix: "/api" } }`)
	if got := extractAutoloadDir(obj, src); got != "" {
		t.Fatalf("expected empty when no dir, got %q", got)
	}
}
