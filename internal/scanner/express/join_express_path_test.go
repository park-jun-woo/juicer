//ff:func feature=scan type=test control=sequence topic=express
//ff:what joinExpressPath: 빈prefix / path루트or빈 / 슬래시 정규화 결합
package express

import "testing"

func TestJoinExpressPath(t *testing.T) {
	cases := []struct{ prefix, path, want string }{
		{"", "/users", "/users"},      // empty prefix
		{"/api", "/", "/api"},         // path "/"
		{"/api", "", "/api"},          // path ""
		{"/api/", "/users", "/api/users"}, // slash normalization
		{"/api", "users", "/api/users"},   // add slash
	}
	for _, c := range cases {
		if got := joinExpressPath(c.prefix, c.path); got != c.want {
			t.Errorf("join(%q,%q)=%q want %q", c.prefix, c.path, got, c.want)
		}
	}
}
