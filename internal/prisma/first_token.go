//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what 선두 공백 토큰과 나머지 문자열 분리
package prisma

// firstToken returns the first whitespace-delimited token and the trimmed rest.
func firstToken(s string) (string, string) {
	i := 0
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}
	start := i
	for i < len(s) && s[i] != ' ' && s[i] != '\t' {
		i++
	}
	tok := s[start:i]
	rest := s[i:]
	return tok, rest
}
