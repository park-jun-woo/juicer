//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what enclosingScopes: 바이트 위치를 포함하는 스코프만 선별함을 검증
package fastify

import "testing"

func TestEnclosingScopes(t *testing.T) {
	scopes := []wrapperScope{
		{Start: 0, End: 100, Prefix: "/a"},
		{Start: 10, End: 50, Prefix: "/b"},
		{Start: 200, End: 300, Prefix: "/c"},
	}
	got := enclosingScopes(20, scopes)
	if len(got) != 2 {
		t.Fatalf("want 2 enclosing, got %d (%v)", len(got), got)
	}
}
