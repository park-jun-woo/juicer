//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 필드의 @map 오버라이드 또는 Prisma 필드명으로 컬럼명 결정
package prisma

// columnName returns the column name, honoring a @map("...") override.
func columnName(f field) string {
	for _, a := range f.attrs {
		if c, ok := mapArg(a, "@map"); ok {
			return c
		}
	}
	return f.name
}
