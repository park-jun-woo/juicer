//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractAutoloadPrefix 테스트
package fastify

import "testing"

func TestExtractAutoloadPrefix_NestedOptions(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "r", options: { prefix: "/api" } }`)
	if got := extractAutoloadPrefix(obj, src); got != "/api" {
		t.Fatalf("nested options prefix = %q, want /api", got)
	}
}

func TestExtractAutoloadPrefix_TopLevel(t *testing.T) {
	// no options.prefix -> falls back to top-level prefix
	obj, src := firstObject(t, `{ dir: "r", prefix: "/v1" }`)
	if got := extractAutoloadPrefix(obj, src); got != "/v1" {
		t.Fatalf("top-level prefix = %q, want /v1", got)
	}
}

func TestExtractAutoloadPrefix_None(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "r" }`)
	if got := extractAutoloadPrefix(obj, src); got != "" {
		t.Fatalf("expected empty prefix, got %q", got)
	}
}
