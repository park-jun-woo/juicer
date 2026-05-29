//ff:func feature=prisma type=parse control=selection topic=prisma
//ff:what 괄호/대괄호 문자에 대한 깊이 증감 반환
package prisma

// attrDepthDelta returns the paren/bracket nesting delta for one byte.
func attrDepthDelta(c byte) int {
	switch c {
	case '(', '[':
		return 1
	case ')', ']':
		return -1
	default:
		return 0
	}
}
