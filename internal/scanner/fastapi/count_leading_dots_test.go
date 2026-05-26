//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what countLeadingDots 테스트
package fastapi

import "testing"

func TestCountLeadingDots(t *testing.T) {
	tests := []struct{ in string; want int }{
		{"", 0}, {".", 1}, {"..", 2}, {"...models", 3}, {"models", 0},
	}
	for _, tt := range tests {
		if got := countLeadingDots(tt.in); got != tt.want {
			t.Errorf("countLeadingDots(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}
