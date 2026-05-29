//ff:func feature=prisma type=parse control=sequence topic=prisma
//ff:what @map/@@map 등 attr의 첫 따옴표 인자 추출
package prisma

import "strings"

// mapArg returns the first quoted argument of the given attribute prefix
// (e.g. @@map("accounts") -> "accounts").
func mapArg(attr, prefix string) (string, bool) {
	if !strings.HasPrefix(attr, prefix) {
		return "", false
	}
	rest := attr[len(prefix):]
	open := strings.IndexByte(rest, '"')
	if open < 0 {
		return "", false
	}
	rest = rest[open+1:]
	close := strings.IndexByte(rest, '"')
	if close < 0 {
		return "", false
	}
	return rest[:close], true
}
