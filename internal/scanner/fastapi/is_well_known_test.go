//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what isWellKnown: 직접/한정 이름 매칭 / 비매칭
package fastapi

import "testing"

func TestIsWellKnown(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"BaseModel", true},
		{"pydantic.BaseModel", true}, // qualified
		{"SQLModel", true},
		{"Unknown", false},
		{"pkg.Unknown", false},
	}
	for _, c := range cases {
		if got := isWellKnown(c.in); got != c.want {
			t.Errorf("isWellKnown(%q)=%v want %v", c.in, got, c.want)
		}
	}
}
