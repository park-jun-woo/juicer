//ff:func feature=scan type=test control=sequence
//ff:what containsAuthKeyword — 인증 키워드 포함 판별을 검증
package scanner

import "testing"

func TestContainsAuthKeyword(t *testing.T) {
	if !containsAuthKeyword("requireAuth") {
		t.Error("expected true for 'requireAuth'")
	}
	if !containsAuthKeyword("JWT_VERIFY") { // case-insensitive
		t.Error("expected true for 'JWT_VERIFY'")
	}
	if containsAuthKeyword("plainHandler") {
		t.Error("expected false for 'plainHandler'")
	}
	if containsAuthKeyword("") {
		t.Error("expected false for empty string")
	}
}
