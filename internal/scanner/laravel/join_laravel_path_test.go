//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what joinLaravelPath 테스트 (round5)
package laravel

import "testing"

func TestJoinLaravelPath_Round5(t *testing.T) {
	cases := []struct {
		prefix, path, want string
	}{
		{"", "", "/"},
		{"", "users", "/users"},
		{"api/", "/users", "/api/users"},
		{"api", "", "/api"},
		{"api/v1", "users/", "/api/v1/users/"},
	}
	for _, c := range cases {
		if got := joinLaravelPath(c.prefix, c.path); got != c.want {
			t.Errorf("joinLaravelPath(%q,%q)=%q want %q", c.prefix, c.path, got, c.want)
		}
	}
}
