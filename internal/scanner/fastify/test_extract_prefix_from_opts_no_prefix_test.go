//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractPrefixFromOpts_NoPrefix 테스트
package fastify

import "testing"

func TestExtractPrefixFromOpts_NoPrefix(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "x" }`)
	if got := extractPrefixFromOpts(obj, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
