//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 테이블 맵 키를 알파벳순 슬라이스로 반환
package ddl

import "sort"

// sortedTableNames returns the map keys in alphabetical order.
func sortedTableNames(tables map[string]*Table) []string {
	names := make([]string, 0, len(tables))
	for name := range tables {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
