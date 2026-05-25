//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestJoinPath 테스트
package scanner

import (
	"testing"
)

func TestJoinPath(t *testing.T) {
	tests := []struct {
		a, b string
		want string
	}{
		{"", "/users", "/users"},
		{"/api", "/users", "/api/users"},
		{"/api/", "/users", "/api/users"},
		{"/api", "users", "/api/users"},
	}
	for _, tt := range tests {
		got := joinPath(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("joinPath(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
		}
	}
}
