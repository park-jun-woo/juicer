//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestNestedOptionsPrefix_OptionsNotObject 테스트
package fastify

import "testing"

func TestNestedOptionsPrefix_OptionsNotObject(t *testing.T) {
	obj, src := firstObject(t, `{ options: "x" }`)
	if got := nestedOptionsPrefix(obj, src); got != "" {
		t.Fatalf("expected empty when options not object, got %q", got)
	}
}
