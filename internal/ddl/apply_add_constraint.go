//ff:func feature=ddl type=parse control=sequence
//ff:what ADD CONSTRAINT 절을 테이블의 Constraints에 추가하고, FK인 경우 이전 이름 없는 FK와 Column.Raw의 REFERENCES를 제거한다
package ddl

import "strings"

// applyAddConstraint appends a constraint definition to the table.
// For FOREIGN KEY constraints, it also:
//  1. Removes any existing unnamed FK for the same column from Constraints.
//  2. Strips the old inline REFERENCES from the target column's Raw.
func applyAddConstraint(t *Table, constraintDef string) {
	def := strings.Join(strings.Fields(constraintDef), " ")

	col := extractFKColumn(def)
	if col != "" {
		removeUnnamedFK(t, col)
	}

	t.Constraints = append(t.Constraints, def)

	if col == "" {
		return
	}
	stripColumnReferences(t, col)
}
