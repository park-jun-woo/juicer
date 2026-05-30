//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what joinHonoPath 테스트
package hono

import "testing"

func TestJoinHonoPath(t *testing.T) {
	cases := []struct{ prefix, path, want string }{
		{"", "/users", "/users"},
		{"/api", "", "/api"},
		{"/api", "/", "/api"},
		{"/api", "/users", "/api/users"},
		{"/api/", "users", "/api/users"},
		{"/api/", "/users", "/api/users"},
	}
	for _, c := range cases {
		if got := joinHonoPath(c.prefix, c.path); got != c.want {
			t.Errorf("joinHonoPath(%q,%q)=%q want %q", c.prefix, c.path, got, c.want)
		}
	}
}
