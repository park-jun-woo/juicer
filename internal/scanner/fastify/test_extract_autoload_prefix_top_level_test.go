//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractAutoloadPrefix_TopLevel 테스트
package fastify

import "testing"

func TestExtractAutoloadPrefix_TopLevel(t *testing.T) {

	obj, src := firstObject(t, `{ dir: "r", prefix: "/v1" }`)
	if got := extractAutoloadPrefix(obj, src); got != "/v1" {
		t.Fatalf("top-level prefix = %q, want /v1", got)
	}
}
