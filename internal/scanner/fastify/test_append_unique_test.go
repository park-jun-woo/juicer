//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what appendUnique: 중복 제거 추가를 검증
package fastify

import "testing"

func TestAppendUnique(t *testing.T) {
	got := appendUnique([]string{"a"}, "a", "b", "b", "c")
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Fatalf("want [a b c], got %v", got)
	}
}
