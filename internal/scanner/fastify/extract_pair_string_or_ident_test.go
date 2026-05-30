//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractPairStringOrIdent 테스트
package fastify

import "testing"

func TestExtractPairStringOrIdent(t *testing.T) {
	obj, src := firstObject(t, `{ s: "hello", id: someVar, n: 42, b: true }`)

	if got := extractPairStringOrIdent(obj, src, "s"); got != "hello" {
		t.Errorf("string: got %q", got)
	}
	if got := extractPairStringOrIdent(obj, src, "id"); got != "someVar" {
		t.Errorf("identifier: got %q", got)
	}
	if got := extractPairStringOrIdent(obj, src, "n"); got != "42" {
		t.Errorf("number: got %q", got)
	}
	// boolean (true) -> default case -> ""
	if got := extractPairStringOrIdent(obj, src, "b"); got != "" {
		t.Errorf("bool default: got %q, want empty", got)
	}
	// missing key -> val nil -> ""
	if got := extractPairStringOrIdent(obj, src, "missing"); got != "" {
		t.Errorf("missing: got %q, want empty", got)
	}
}
