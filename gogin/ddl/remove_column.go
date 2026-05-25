//ff:func feature=ddl type=extract control=iteration dimension=1
//ff:what 컬럼 슬라이스에서 이름으로 컬럼 제거
package ddl

// removeColumn removes a column by name from the columns slice.
func removeColumn(cols []Column, name string) []Column {
	result := make([]Column, 0, len(cols))
	for _, c := range cols {
		if c.Name != name {
			result = append(result, c)
		}
	}
	return result
}
