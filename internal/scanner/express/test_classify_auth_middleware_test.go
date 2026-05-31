//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 단위 테스트: classifyAuthMiddleware 분류 결과를 검증한다
package express

import "testing"

func TestExtractAuthMiddleware_ClassifyAuthMiddleware(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"auth", "auth"},
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
		// Phase139: 부분일치 — member_expression 전체 텍스트(객체 한정자 포함).
		{"mw.authAdminApi", "auth"},
		{"mw.authAdminApiWithUrl", "auth"},
		{"mw.authenticatePublic", "auth"},
		{"middleware.authMemberByUuid", "auth"},
		{"ensureLoggedIn", "auth"},
		{"JwtAuthGuard", "auth"},
		{"AdminGuard", "auth"},
		// public 유지 (auth 부분문자열 없음 / auth 어근 없는 헬퍼).
		{"mw.publicAdminApi", ""},
		{"requireAll", ""},
		{"ensureDir", ""},
	}
	for _, tt := range tests {
		got := classifyAuthMiddleware(tt.name)
		if got != tt.want {
			t.Errorf("classifyAuthMiddleware(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}
