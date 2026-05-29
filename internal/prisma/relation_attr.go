//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what 속성 목록에서 @relation(...) 내부 인자 문자열 추출
package prisma

import "strings"

// relationAttr returns the inner argument string of an @relation(...) attribute.
func relationAttr(attrs []string) (string, bool) {
	for _, a := range attrs {
		if !strings.HasPrefix(a, "@relation(") {
			continue
		}
		rest := a[len("@relation("):]
		if i := strings.LastIndexByte(rest, ')'); i >= 0 {
			rest = rest[:i]
		}
		return rest, true
	}
	return "", false
}
