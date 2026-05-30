//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractPrefixFromOpts_Template 테스트
package fastify

import "testing"

func TestExtractPrefixFromOpts_Template(t *testing.T) {
	obj, src := firstObject(t, "{ prefix: `/v2` }")
	if got := extractPrefixFromOpts(obj, src); got != "/v2" {
		t.Fatalf("template prefix: got %q", got)
	}
}
