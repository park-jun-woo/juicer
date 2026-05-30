//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractParamName 테스트
package hono

import "testing"

func TestExtractParamName(t *testing.T) {
	cases := []struct{ in, want string }{
		{":id", "id"},
		{"users", ""},
		{":id{[0-9]+}", "id{[0-9]+}"}, // no paren, kept as-is
		{":id(\\d+)", "id"},           // regex paren stripped
		{"", ""},
		{":", ""},
	}
	for _, c := range cases {
		if got := extractParamName(c.in); got != c.want {
			t.Errorf("extractParamName(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
