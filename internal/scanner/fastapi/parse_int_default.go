//ff:func feature=scan type=parse control=iteration dimension=1 topic=fastapi
//ff:what 문자열을 정수로 파싱하여 실패 시 기본값을 반환한다
package fastapi

// parseIntDefault parses a string as int, returning def on failure.
func parseIntDefault(s string, def int) int {
	if len(s) == 0 {
		return def
	}
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return def
		}
		n = n*10 + int(ch-'0')
	}
	return n
}
