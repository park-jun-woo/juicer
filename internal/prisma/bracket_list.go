//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what 문자열 첫 [ ... ] 안의 콤마 구분 토큰 목록 추출
package prisma

import "strings"

// bracketList returns the comma-separated tokens inside the first [...] of s.
func bracketList(s string) []string {
	open := strings.IndexByte(s, '[')
	if open < 0 {
		return nil
	}
	rest := s[open+1:]
	close := strings.IndexByte(rest, ']')
	if close < 0 {
		return nil
	}
	parts := strings.Split(rest[:close], ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	return out
}
