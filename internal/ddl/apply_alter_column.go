//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what ALTER COLUMN 절을 대상 컬럼의 Raw에 적용한다
package ddl

import "strings"

// applyAlterColumn applies an ALTER COLUMN action to the named column.
func applyAlterColumn(t *Table, colName, action string) {
	colName = strings.ToLower(colName)
	for i, c := range t.Columns {
		if c.Name == colName {
			t.Columns[i].Raw = modifyColumnRaw(c.Raw, action)
			return
		}
	}
}
