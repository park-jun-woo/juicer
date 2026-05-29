//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 모델의 필드명->컬럼명(@map 반영) 매핑 생성
package prisma

// fieldColumnMap maps each field name in m to its column name (honoring @map).
func fieldColumnMap(m model) map[string]string {
	out := make(map[string]string, len(m.fields))
	for _, f := range m.fields {
		out[f.name] = columnName(f)
	}
	return out
}
