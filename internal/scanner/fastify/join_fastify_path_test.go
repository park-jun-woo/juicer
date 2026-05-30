//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what joinFastifyPath 테스트
package fastify

import "testing"

func TestJoinFastifyPath(t *testing.T) {
	tests := []struct{ prefix, path, want string }{
		{"", "/users", "/users"},        // empty prefix
		{"/api", "/", "/api"},           // root path
		{"/api", "", "/api"},            // empty path
		{"/api", "/users", "/api/users"},
		{"/api/", "users", "/api/users"}, // slash dedup
	}
	for _, tt := range tests {
		if got := joinFastifyPath(tt.prefix, tt.path); got != tt.want {
			t.Errorf("joinFastifyPath(%q,%q) = %q, want %q", tt.prefix, tt.path, got, tt.want)
		}
	}
}
