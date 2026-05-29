//ff:func feature=prisma type=parse control=sequence topic=prisma
//ff:what name: [a, b] 형태의 named 인자에서 토큰 목록 추출
package prisma

import "strings"

// namedBracketList returns the bracket token list of a named argument
// (e.g. fields: [authorId, orgId]) within s.
func namedBracketList(s, name string) []string {
	idx := strings.Index(s, name)
	if idx < 0 {
		return nil
	}
	return bracketList(s[idx:])
}
