//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what bracketDelta 테스트
package fastapi

import "testing"

func TestBracketDelta(t *testing.T) {
	tests := []struct {
		ch   rune
		want int
	}{
		{'[', 1}, {'(', 1}, {']', -1}, {')', -1}, {'a', 0}, {' ', 0},
	}
	for _, tt := range tests {
		if got := bracketDelta(tt.ch); got != tt.want {
			t.Errorf("bracketDelta(%q) = %d, want %d", tt.ch, got, tt.want)
		}
	}
}
