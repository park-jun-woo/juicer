//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what wrapperPrefixFor: 바깥→안쪽 스코프 prefix 합성과 미포함 시 빈 문자열을 검증
package fastify

import "testing"

func TestWrapperPrefixFor(t *testing.T) {
	scopes := []wrapperScope{
		{Start: 0, End: 100, Prefix: "/outer"},
		{Start: 10, End: 50, Prefix: "/inner"},
	}
	if got := wrapperPrefixFor(20, scopes); got != "/outer/inner" {
		t.Fatalf("want /outer/inner, got %q", got)
	}
	if got := wrapperPrefixFor(500, scopes); got != "" {
		t.Fatalf("want empty, got %q", got)
	}
}
