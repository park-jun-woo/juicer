//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what parseStatusCode 테스트
package nestjs

import "testing"

func TestParseStatusCode_Cases(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"200", 200},
		{"201", 201},
		{"404", 404},
		{"", 0},
		{"abc", 0},
		{"2xx", 2},
	}
	for _, c := range cases {
		got := parseStatusCode(c.in)
		if got != c.want {
			t.Errorf("parseStatusCode(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
