//ff:func feature=ddl type=parse control=sequence
//ff:what 단일 ADD COLUMN 또는 DROP COLUMN 절을 테이블에 적용한다
package ddl

import "strings"

// applyAlterClause applies a single ADD COLUMN or DROP COLUMN clause to the table.
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
	}
}
