//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 문자가 Python 식별자의 유효한 문자인지 확인한다
package fastapi

// isIdentChar returns true if ch is valid at position in a Python identifier.
// firstPos=true means the character must be a letter or underscore.
func isIdentChar(ch rune, firstPos bool) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' {
		return true
	}
	if !firstPos && ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
