//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what 모델의 PK/UNIQUE/FK 제약 라인들을 조립
package prisma

// buildConstraints assembles PRIMARY KEY / UNIQUE / FOREIGN KEY lines for m.
func buildConstraints(m model, s schema) []string {
	cons := make([]string, 0, 4)
	cons = appendIfNotEmpty(cons, primaryKeyConstraint(m, s))
	cons = append(cons, uniqueConstraints(m, s)...)
	cons = append(cons, foreignKeyConstraints(m, s)...)
	return cons
}
