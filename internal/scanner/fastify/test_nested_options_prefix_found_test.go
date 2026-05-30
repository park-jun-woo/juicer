//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestNestedOptionsPrefix_Found 테스트
package fastify

import "testing"

func TestNestedOptionsPrefix_Found(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "r", options: { prefix: "/api" } }`)
	if got := nestedOptionsPrefix(obj, src); got != "/api" {
		t.Fatalf("got %q, want /api", got)
	}
}
