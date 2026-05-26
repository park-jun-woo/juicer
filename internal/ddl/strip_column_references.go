//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what 테이블 컬럼에서 지정 컬럼의 inline REFERENCES를 제거한다
package ddl

import "strings"

// stripColumnReferences finds a column by name in the table and removes the
// inline REFERENCES clause from its Raw definition.
func stripColumnReferences(t *Table, colName string) {
	colLower := strings.ToLower(colName)
	for i := range t.Columns {
		if t.Columns[i].Name == colLower {
			t.Columns[i].Raw = stripInlineReferences(t.Columns[i].Raw)
			break
		}
	}
}
