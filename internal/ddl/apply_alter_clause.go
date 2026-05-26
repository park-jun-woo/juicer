//ff:func feature=ddl type=parse control=sequence
//ff:what 단일 ALTER TABLE 절을 테이블에 적용한다 (ADD/DROP COLUMN, ALTER COLUMN, ADD/DROP CONSTRAINT)
package ddl

import "strings"

// applyAlterClause applies a single ALTER TABLE sub-clause to the table.
func applyAlterClause(t *Table, clause string) {
	clause = strings.TrimSpace(clause)
	if clause == "" {
		return
	}
	if m := reAddColumn.FindStringSubmatch(clause); m != nil {
		colDef := strings.TrimSpace(m[1])
		colName := extractColumnName(colDef)
		if colName != "" && !hasColumn(t, colName) {
			t.Columns = append(t.Columns, Column{Name: colName, Raw: colDef})
		}
	} else if m := reDropColumn.FindStringSubmatch(clause); m != nil {
		colName := strings.ToLower(m[1])
		t.Columns = removeColumn(t.Columns, colName)
	} else if m := reAlterColumn.FindStringSubmatch(clause); m != nil {
		applyAlterColumn(t, m[1], m[2])
	} else if m := reAddConstraint.FindStringSubmatch(clause); m != nil {
		applyAddConstraint(t, "CONSTRAINT "+m[1])
	} else if m := reDropConstraint.FindStringSubmatch(clause); m != nil {
		applyDropConstraint(t, m[1])
	}
}
