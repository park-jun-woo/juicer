//ff:func feature=ddl type=parse control=sequence
//ff:what ADD CONSTRAINT 절을 테이블의 Constraints에 추가한다
package ddl

import "strings"

// applyAddConstraint appends a constraint definition to the table.
func applyAddConstraint(t *Table, constraintDef string) {
	def := strings.Join(strings.Fields(constraintDef), " ")
	t.Constraints = append(t.Constraints, def)
}
