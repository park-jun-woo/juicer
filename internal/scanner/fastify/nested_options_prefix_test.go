//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what nestedOptionsPrefix 테스트
package fastify

import "testing"

func TestNestedOptionsPrefix_Found(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "r", options: { prefix: "/api" } }`)
	if got := nestedOptionsPrefix(obj, src); got != "/api" {
		t.Fatalf("got %q, want /api", got)
	}
}

func TestNestedOptionsPrefix_NoOptions(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "r" }`)
	if got := nestedOptionsPrefix(obj, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestNestedOptionsPrefix_OptionsNotObject(t *testing.T) {
	obj, src := firstObject(t, `{ options: "x" }`)
	if got := nestedOptionsPrefix(obj, src); got != "" {
		t.Fatalf("expected empty when options not object, got %q", got)
	}
}
