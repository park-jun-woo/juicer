//ff:func feature=scan type=test control=sequence topic=actix
//ff:what hasPrefix — 경로 prefix 판별의 분기를 검증
package actix

import "testing"

func TestHasPrefix(t *testing.T) {
	cases := []struct {
		path, prefix string
		want         bool
	}{
		{"/api/users", "", true},      // empty prefix -> true
		{"/api/users", "/", true},     // root prefix -> true
		{"/api/users", "/api", true},  // real prefix match
		{"/api/users", "/x", false},   // mismatch
		{"/ap", "/api", false},        // shorter than prefix
		{"/api", "/api", true},        // exact length match
	}
	for _, c := range cases {
		if got := hasPrefix(c.path, c.prefix); got != c.want {
			t.Errorf("hasPrefix(%q, %q) = %v, want %v", c.path, c.prefix, got, c.want)
		}
	}
}
