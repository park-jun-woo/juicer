//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what pathToOpenAPI 테스트
package nestjs

import "testing"

func TestPathToOpenAPI_Cases(t *testing.T) {
	cases := []struct{ in, want string }{
		{"/users/:id", "/users/{id}"},
		{"/users", "/users"},
		{"/users/:id/posts/:pid", "/users/{id}/posts/{pid}"},
		{"/:id", "/{id}"},
	}
	for _, c := range cases {
		got := pathToOpenAPI(c.in)
		if got != c.want {
			t.Errorf("pathToOpenAPI(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
