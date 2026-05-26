//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestIsAuthMiddleware 테스트
package fastapi

import "testing"

func TestIsAuthMiddleware(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"get_current_user", true},
		{"auth_required", true},
		{"current_user_dep", true},
		{"validate_data", false},
	}
	for _, tc := range tests {
		got := isAuthMiddleware(tc.name)
		if got != tc.want {
			t.Errorf("isAuthMiddleware(%q) = %v, want %v", tc.name, got, tc.want)
		}
	}
}
