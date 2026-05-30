//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestNestedOptionsPrefix_NoOptions 테스트
package fastify

import "testing"

func TestNestedOptionsPrefix_NoOptions(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "r" }`)
	if got := nestedOptionsPrefix(obj, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
