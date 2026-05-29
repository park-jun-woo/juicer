//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what 블록 속성(@@id/@@unique/@@index)의 [a, b] 필드 이름 목록 추출
package prisma

import "strings"

// blockKeyFields returns the field-name list inside the first block attribute
// with the given prefix (e.g. @@unique([orgId, email]) -> [orgId email]).
func blockKeyFields(attrs []string, prefix string) ([]string, bool) {
	for _, a := range attrs {
		if !strings.HasPrefix(a, prefix) {
			continue
		}
		return bracketList(a), true
	}
	return nil, false
}
