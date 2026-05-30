//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what isIdentChar: 문자/언더스코어 / 숫자(첫자리 제외) / 무효
package fastapi

import "testing"

func TestIsIdentChar(t *testing.T) {
	cases := []struct {
		ch       rune
		firstPos bool
		want     bool
	}{
		{'a', true, true},
		{'Z', false, true},
		{'_', true, true},
		{'5', true, false},  // digit not allowed at first position
		{'5', false, true},  // digit allowed elsewhere
		{'-', false, false}, // invalid char
		{'$', true, false},
	}
	for _, c := range cases {
		if got := isIdentChar(c.ch, c.firstPos); got != c.want {
			t.Errorf("isIdentChar(%q,%v)=%v want %v", c.ch, c.firstPos, got, c.want)
		}
	}
}
