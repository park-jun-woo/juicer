//ff:func feature=scan type=parse control=iteration dimension=1 topic=nestjs
//ff:what 문자열에서 상태코드 숫자를 파싱한다
package nestjs

// parseStatusCode converts string to int.
func parseStatusCode(s string) int {
	code := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			code = code*10 + int(c-'0')
		}
	}
	return code
}
