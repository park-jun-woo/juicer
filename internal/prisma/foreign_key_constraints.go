//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what @relation 필드로부터 FOREIGN KEY ... REFERENCES ... 제약 라인들 생성
package prisma

// foreignKeyConstraints returns FOREIGN KEY lines from @relation fields that
// declare fields/references (the owning side of the relation).
func foreignKeyConstraints(m model, s schema) []string {
	out := make([]string, 0, 2)
	for _, f := range m.fields {
		rel, ok := relationAttr(f.attrs)
		if !ok {
			continue
		}
		out = appendIfNotEmpty(out, relationLine(m, f, rel, s))
	}
	return out
}
