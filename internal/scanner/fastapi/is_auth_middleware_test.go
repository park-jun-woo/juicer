//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what isAuthMiddleware 테스트
package fastapi

import "testing"

func TestIsAuthMiddleware_All(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"get_current_user", true},
		{"auth_required", true},
		{"get_current_active_user", true},
		{"get_db", false},
		{"parse_data", false},
	}
	for _, tt := range tests {
		if got := isAuthMiddleware(tt.in); got != tt.want {
			t.Errorf("isAuthMiddleware(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
