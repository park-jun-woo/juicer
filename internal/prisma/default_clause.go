//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 필드 @default(...)를 SQL DEFAULT 식으로 변환 (autoincrement은 빈 식)
package prisma

// defaultClause returns the SQL DEFAULT expression for a field's @default(...).
// autoincrement() yields ("", true) so serial promotion handles it without a
// DEFAULT clause. The bool reports whether a @default attribute was present.
func defaultClause(f field) (string, bool) {
	for _, a := range f.attrs {
		inner, ok := defaultArg(a)
		if !ok {
			continue
		}
		return mapDefaultExpr(inner), true
	}
	return "", false
}
