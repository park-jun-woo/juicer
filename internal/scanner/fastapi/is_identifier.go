//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 문자열이 유효한 Python 식별자인지 확인한다
package fastapi

// isIdentifier returns true if s is a valid Python identifier (starts with
// letter/underscore, contains only letters, digits, underscores).
func isIdentifier(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, ch := range s {
		if !isIdentChar(ch, i == 0) {
			return false
		}
	}
	return true
}
