//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what composePrefixes: 부모 없음/다중 부모 합성을 검증
package fastify

import "testing"

func TestComposePrefixes(t *testing.T) {
	if got := composePrefixes(nil, "/x"); len(got) != 1 || got[0] != "/x" {
		t.Fatalf("empty parents: want [/x], got %v", got)
	}
	got := composePrefixes([]string{"/a", "/b"}, "/c")
	if len(got) != 2 || got[0] != "/a/c" || got[1] != "/b/c" {
		t.Fatalf("want [/a/c /b/c], got %v", got)
	}
}
