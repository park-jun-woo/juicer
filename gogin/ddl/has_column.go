//ff:func feature=ddl type=extract control=iteration dimension=1
//ff:what 테이블에 해당 컬럼이 있는지 확인
package ddl

// hasColumn checks if the table already has a column with the given name.
func hasColumn(t *Table, name string) bool {
	for _, c := range t.Columns {
		if c.Name == name {
			return true
		}
	}
	return false
}
