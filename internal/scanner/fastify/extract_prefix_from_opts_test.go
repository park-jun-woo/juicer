//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractPrefixFromOpts 테스트
package fastify

import "testing"

func TestExtractPrefixFromOpts_String(t *testing.T) {
	obj, src := firstObject(t, `{ prefix: "/api" }`)
	if got := extractPrefixFromOpts(obj, src); got != "/api" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractPrefixFromOpts_Template(t *testing.T) {
	obj, src := firstObject(t, "{ prefix: `/v2` }")
	if got := extractPrefixFromOpts(obj, src); got != "/v2" {
		t.Fatalf("template prefix: got %q", got)
	}
}

func TestExtractPrefixFromOpts_NoPrefix(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "x" }`)
	if got := extractPrefixFromOpts(obj, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestExtractPrefixFromOpts_NonString(t *testing.T) {
	// prefix value is not a string/template -> ""
	obj, src := firstObject(t, `{ prefix: someVar }`)
	if got := extractPrefixFromOpts(obj, src); got != "" {
		t.Fatalf("expected empty for non-string prefix, got %q", got)
	}
}
