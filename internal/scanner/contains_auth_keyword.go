//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 문자열에 인증 관련 키워드가 포함되어 있는지 확인한다
package scanner

import "strings"

var authKeywords = []string{"auth", "guard", "jwt", "current_user", "get_current", "oauth", "token", "verify"}

func containsAuthKeyword(s string) bool {
	lower := strings.ToLower(s)
	for _, kw := range authKeywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	return false
}
