//ff:func feature=scan type=parse control=selection topic=fastapi
//ff:what 괄호 문자에 대한 깊이 변화값을 반환한다
package fastapi

// bracketDelta returns +1 for opening brackets, -1 for closing, 0 otherwise.
func bracketDelta(ch rune) int {
	switch ch {
	case '[', '(':
		return 1
	case ']', ')':
		return -1
	default:
		return 0
	}
}
