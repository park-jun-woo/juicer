//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractAutoloadPrefix_NestedOptions 테스트
package fastify

import "testing"

func TestExtractAutoloadPrefix_NestedOptions(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "r", options: { prefix: "/api" } }`)
	if got := extractAutoloadPrefix(obj, src); got != "/api" {
		t.Fatalf("nested options prefix = %q, want /api", got)
	}
}
