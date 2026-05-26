//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what joinPath 테스트
package fastapi

import "testing"

func TestJoinPath(t *testing.T) {
	tests := []struct {
		parts []string
		want  string
	}{
		{[]string{"/api", "/users"}, "/api/users"},
		{[]string{"", ""}, ""},
		{[]string{"/api/", "/v1/"}, "/api/v1"},
		{[]string{"api", "users", "list"}, "/api/users/list"},
	}
	for _, tt := range tests {
		got := joinPath(tt.parts...)
		if got != tt.want {
			t.Errorf("joinPath(%v) = %q, want %q", tt.parts, got, tt.want)
		}
	}
}
