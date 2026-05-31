//ff:func feature=prisma type=test control=iteration dimension=1 topic=prisma
//ff:what attrDepthDelta 괄호/대괄호 깊이 증감 테스트
package prisma

import "testing"

func TestAttrDepthDelta(t *testing.T) {
	cases := map[byte]int{'(': 1, '[': 1, ')': -1, ']': -1, 'a': 0, ' ': 0}
	for c, want := range cases {
		if got := attrDepthDelta(c); got != want {
			t.Errorf("attrDepthDelta(%q) = %d, want %d", c, got, want)
		}
	}
}
