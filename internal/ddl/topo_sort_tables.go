//ff:func feature=ddl type=render control=sequence
//ff:what FK 의존 그래프로 테이블명을 위상정렬 (참조 대상 먼저, 동률 알파벳순, 순환 시 경고+fallback)
package ddl

import "regexp"

// reFKReference captures the referenced table name from a FOREIGN KEY clause.
// The target may be quoted ("User"), schema-qualified (public.users) or plain.
var reFKReference = regexp.MustCompile(`(?i)REFERENCES\s+` + tblName)

// topoSortTables returns table-map keys ordered so that a referenced table
// always precedes any table that references it (via a FOREIGN KEY constraint).
// Tables with no dependency relation, and the members of any reference cycle,
// fall back to alphabetical order so the result is deterministic.
func topoSortTables(tables map[string]*Table) []string {
	names := sortedTableNames(tables)
	deps := buildFKDeps(tables, names)
	return kahnSort(names, deps)
}
