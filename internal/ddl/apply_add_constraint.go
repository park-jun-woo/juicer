//ff:func feature=ddl type=parse control=sequence
//ff:what ADD CONSTRAINT 절을 테이블의 Constraints에 추가하고, FK인 경우 Column.Raw의 이전 REFERENCES를 제거한다
package ddl

import "strings"

// applyAddConstraint appends a constraint definition to the table.
// For FOREIGN KEY constraints, it also strips the old inline REFERENCES from the
// target column's Raw so that render_table does not emit a stale FK reference.
func applyAddConstraint(t *Table, constraintDef string) {
	def := strings.Join(strings.Fields(constraintDef), " ")
	t.Constraints = append(t.Constraints, def)

	col := extractFKColumn(def)
	if col == "" {
		return
	}
	stripColumnReferences(t, col)
}
