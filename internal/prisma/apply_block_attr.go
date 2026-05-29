//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what @@ 블록 속성 라인을 model에 기록하고 @@map이면 테이블명 갱신
package prisma

// applyBlockAttr records a @@ block attribute line on m, updating the table
// name when the line is an @@map override.
func applyBlockAttr(m *model, line string) {
	m.blockAttrs = append(m.blockAttrs, line)
	if t, ok := mapArg(line, "@@map"); ok {
		m.tableName = t
	}
}
