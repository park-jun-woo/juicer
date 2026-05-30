//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractPrefixFromOpts_String 테스트
package fastify

import "testing"

func TestExtractPrefixFromOpts_String(t *testing.T) {
	obj, src := firstObject(t, `{ prefix: "/api" }`)
	if got := extractPrefixFromOpts(obj, src); got != "/api" {
		t.Fatalf("got %q", got)
	}
}
