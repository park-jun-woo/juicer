//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 테이블 상태를 최종 DDL 문자열로 렌더링
package ddl

import (
	"strings"
)

// Render produces the final DDL output from enum types and the table state map.
// Enum types are emitted (sorted) before all tables; tables are emitted in
// FK-dependency topological order so referenced tables precede their referrers.
func Render(enums []EnumType, tables map[string]*Table) string {
	names := topoSortTables(tables)

	var sb strings.Builder

	sorted := sortedEnums(enums)
	for _, e := range sorted {
		renderEnum(&sb, e)
	}
	if len(sorted) > 0 && len(names) > 0 {
		sb.WriteByte('\n')
	}

	for i, name := range names {
		if i > 0 {
			sb.WriteByte('\n')
		}
		t := tables[name]
		renderTable(&sb, t)
	}
	return sb.String()
}
