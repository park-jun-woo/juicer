//ff:func feature=prisma type=parse control=sequence topic=prisma
//ff:what @default(...) 속성의 괄호 내부 식 추출
package prisma

import "strings"

// defaultArg extracts the inner text of an @default(...) attribute.
func defaultArg(attr string) (string, bool) {
	if !strings.HasPrefix(attr, "@default(") {
		return "", false
	}
	rest := attr[len("@default("):]
	if i := strings.LastIndexByte(rest, ')'); i >= 0 {
		rest = rest[:i]
	}
	return strings.TrimSpace(rest), true
}
