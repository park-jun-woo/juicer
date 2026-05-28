//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 단위 테스트: classifyAuthMiddleware 분류 결과를 검증한다
package express

import "testing"

func TestExtractAuthMiddleware_ClassifyAuthMiddleware(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"authenticate", "auth"},
		{"requireAuth", "auth"},
		{"ensureAuthenticated", "auth"},
		{"isAuthenticated", "auth"},
		{"authMiddleware", "auth"},
		{"jwtAuth", "auth"},
		{"verifyToken", "auth"},
		{"passport.authenticate", "auth"},
		{"authorize", "role"},
		{"requireRole", "role"},
		{"checkRole", "role"},
		{"hasRole", "role"},
		{"allowRoles", "role"},
		{"rbac", "role"},
		{"validateRequest", ""},
		{"logger", ""},
		{"cors", ""},
	}
	for _, tt := range tests {
		got := classifyAuthMiddleware(tt.name)
		if got != tt.want {
			t.Errorf("classifyAuthMiddleware(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}
