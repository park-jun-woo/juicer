//ff:func feature=scan type=parse control=iteration dimension=1 topic=nestjs
//ff:what 문자열에서 상태코드 숫자를 파싱한다
package nestjs

import "strings"

// parseStatusCode converts a status code string to int.
// It handles both numeric literals ("200") and HttpStatus member expressions
// ("HttpStatus.OK") by looking up the property name in httpStatusMap.
func parseStatusCode(s string) int {
	if idx := strings.LastIndex(s, "."); idx >= 0 {
		prop := s[idx+1:]
		if code, ok := httpStatusMap[prop]; ok {
			return code
		}
		return 0
	}
	code := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			code = code*10 + int(c-'0')
		}
	}
	return code
}
