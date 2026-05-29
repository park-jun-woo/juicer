//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 속성 목록에 특정 @속성이 존재하는지 검사
package prisma

import "strings"

// hasAttr reports whether attrs contains the given @attribute name exactly
// (e.g. @id, @unique), ignoring those with extra suffixes like @id2.
func hasAttr(attrs []string, name string) bool {
	for _, a := range attrs {
		if a == name {
			return true
		}
		if strings.HasPrefix(a, name+"(") {
			return true
		}
	}
	return false
}
