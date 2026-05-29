//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 필드 @unique 및 블록 @@unique([...])로부터 UNIQUE 제약 라인들 생성
package prisma

import "strings"

// uniqueConstraints returns UNIQUE lines from field @unique attributes and
// every block @@unique([...]) (composite) attribute.
func uniqueConstraints(m model, s schema) []string {
	out := make([]string, 0, 2)
	for _, f := range m.fields {
		if hasAttr(f.attrs, "@unique") {
			out = append(out, "UNIQUE ("+columnName(f)+")")
		}
	}
	for _, a := range m.blockAttrs {
		if !strings.HasPrefix(a, "@@unique") {
			continue
		}
		cols := resolveColumns(m.name, bracketList(a), s)
		out = append(out, "UNIQUE ("+strings.Join(cols, ", ")+")")
	}
	return out
}
