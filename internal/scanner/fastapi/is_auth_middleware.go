//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 의존성 함수명이 인증 관련인지 휴리스틱으로 판별한다
package fastapi

import "strings"

// isAuthMiddleware heuristically determines if a dependency function is auth-related.
func isAuthMiddleware(funcName string) bool {
	lower := strings.ToLower(funcName)
	for _, kw := range dependsKeywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	return false
}
