//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what prefixWidth 숫자 접두 자리수(최소 2) 테스트
package ddl

import "testing"

func TestPrefixWidth(t *testing.T) {
	cases := map[int]int{0: 2, 5: 2, 9: 2, 10: 2, 99: 2, 100: 3, 999: 3, 1000: 4}
	for n, want := range cases {
		if got := prefixWidth(n); got != want {
			t.Errorf("prefixWidth(%d) = %d, want %d", n, got, want)
		}
	}
}
