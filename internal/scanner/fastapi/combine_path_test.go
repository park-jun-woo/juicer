//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what combinePath 테스트
package fastapi

import "testing"

func TestCombinePath(t *testing.T) {
	tests := []struct {
		prefix, path, want string
	}{
		{"/users", "/list", "/users/list"},
		{"/users/", "/list", "/users/list"},
		{"/users", "", "/users"},
		{"/users", "/", "/users"},
		{"", "/", "/"},
		{"", "", "/"},
		{"/api", "items", "/api/items"},
	}
	for _, tt := range tests {
		got := combinePath(tt.prefix, tt.path)
		if got != tt.want {
			t.Errorf("combinePath(%q, %q) = %q, want %q", tt.prefix, tt.path, got, tt.want)
		}
	}
}
