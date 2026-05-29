//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what 선두 단어를 공백/콤마/괄호 경계까지 추출
package prisma

// firstWord returns the leading word terminated by space, comma, bracket or
// paren, plus the remainder of the string.
func firstWord(s string) (string, string) {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ', '\t', ',', ')', ']':
			return s[:i], s[i:]
		}
	}
	return s, ""
}
