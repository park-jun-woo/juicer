//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what isDottedIdentifier: ident.ident true / 점없음 / 무효 부분 false
package fastapi

import "testing"

func TestIsDottedIdentifier(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"items.router", true},
		{"router", false},          // no dot
		{"items.123", false},       // second part invalid
		{"1bad.router", false},     // first part invalid
		{"items.sub.router", false}, // SplitN(2) -> parts[1]="sub.router" not a plain ident
	}
	for _, c := range cases {
		got := isDottedIdentifier(c.in)
		if got != c.want {
			t.Errorf("isDottedIdentifier(%q)=%v want %v", c.in, got, c.want)
		}
	}
}
