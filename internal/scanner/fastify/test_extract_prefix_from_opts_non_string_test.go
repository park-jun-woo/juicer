//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractPrefixFromOpts_NonString 테스트
package fastify

import "testing"

func TestExtractPrefixFromOpts_NonString(t *testing.T) {

	obj, src := firstObject(t, `{ prefix: someVar }`)
	if got := extractPrefixFromOpts(obj, src); got != "" {
		t.Fatalf("expected empty for non-string prefix, got %q", got)
	}
}
