//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 필드 @id 또는 블록 @@id로부터 PRIMARY KEY 제약 라인 생성
package prisma

import "strings"

// primaryKeyConstraint returns the PRIMARY KEY line from a field @id or a
// block @@id([...]), or "" when the model has no primary key.
func primaryKeyConstraint(m model, s schema) string {
	if fieldNames, ok := blockKeyFields(m.blockAttrs, "@@id"); ok {
		return "PRIMARY KEY (" + strings.Join(resolveColumns(m.name, fieldNames, s), ", ") + ")"
	}
	for _, f := range m.fields {
		if hasAttr(f.attrs, "@id") {
			return "PRIMARY KEY (" + columnName(f) + ")"
		}
	}
	return ""
}
