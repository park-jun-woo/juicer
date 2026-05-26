//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what needsResponseDTO 테스트
package nestjs

import "testing"

func TestNeedsResponseDTO_Cases(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"any", false},
		{"string", false},
		{"number", false},
		{"MyDto", true},
		{"void", false},
	}
	for _, c := range cases {
		got := needsResponseDTO(c.in)
		if got != c.want {
			t.Errorf("needsResponseDTO(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
