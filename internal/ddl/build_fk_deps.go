//ff:func feature=ddl type=render control=iteration dimension=2
//ff:what 각 테이블 키를 그것이 참조하는 테이블 키 집합으로 매핑 (외부 참조 제외)
package ddl

// buildFKDeps maps each table key to the set of table keys it references.
// External references (not present in the table map) are dropped.
func buildFKDeps(tables map[string]*Table, names []string) map[string]map[string]bool {
	deps := make(map[string]map[string]bool, len(names))
	for _, name := range names {
		set := make(map[string]bool)
		for _, con := range tables[name].Constraints {
			ref := resolveRefKey(tables, con)
			if ref != "" && ref != name {
				set[ref] = true
			}
		}
		deps[name] = set
	}
	return deps
}
