//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 테이블 상태를 최종 DDL 문자열로 렌더링
package ddl

import (
	"sort"
	"strings"
)

// Render produces the final DDL output from the table state map.
// Tables are sorted alphabetically.
func Render(tables map[string]*Table) string {
	names := make([]string, 0, len(tables))
	for name := range tables {
		names = append(names, name)
	}
	sort.Strings(names)

	var sb strings.Builder
	for i, name := range names {
		if i > 0 {
			sb.WriteByte('\n')
		}
		t := tables[name]
		renderTable(&sb, t)
	}
	return sb.String()
}
