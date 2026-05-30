//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractAutoloadDir 테스트
package fastify

import "testing"

func TestExtractAutoloadDir_StringLiteral(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "routes" }`)
	if got := extractAutoloadDir(obj, src); got != "routes" {
		t.Fatalf("string dir = %q", got)
	}
}

func TestExtractAutoloadDir_JoinCall(t *testing.T) {
	obj, src := firstObject(t, `{ dir: join(__dirname, "routes", "api") }`)
	if got := extractAutoloadDir(obj, src); got != "routes/api" {
		t.Fatalf("join dir = %q, want routes/api", got)
	}
}

func TestExtractAutoloadDir_NoDir(t *testing.T) {
	obj, src := firstObject(t, `{ options: { prefix: "/api" } }`)
	if got := extractAutoloadDir(obj, src); got != "" {
		t.Fatalf("expected empty when no dir, got %q", got)
	}
}

func TestExtractAutoloadDir_EmptySegments(t *testing.T) {
	// dir value is a call with no string args -> no segments -> ""
	obj, src := firstObject(t, `{ dir: resolve(base) }`)
	if got := extractAutoloadDir(obj, src); got != "" {
		t.Fatalf("expected empty for no string segments, got %q", got)
	}
}
