//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what ALTER TABLE 하위 절 처리 (ADD/DROP COLUMN, RENAME TO)
package ddl

import "strings"

// applyAlterTable handles ALTER TABLE sub-clauses (ADD COLUMN, DROP COLUMN, RENAME TO).
// A single ALTER TABLE may contain multiple sub-clauses separated by commas.
func applyAlterTable(tables map[string]*Table, name, rest string) {
	name = strings.ToLower(name)
	t := tables[name]

	// Handle RENAME TO
	if m := reRenameTable.FindStringSubmatch(rest); m != nil {
		if t != nil {
			newName := strings.ToLower(m[1])
			delete(tables, name)
			t.Name = newName
			tables[newName] = t
		}
		return
	}

	// Skip ALTER TABLE statements that don't modify columns (ALTER COLUMN, ADD CONSTRAINT, DROP CONSTRAINT, etc.)
	restUpper := strings.ToUpper(strings.TrimSpace(rest))
	if !strings.Contains(restUpper, "ADD COLUMN") && !strings.Contains(restUpper, "DROP COLUMN") {
		return
	}

	if t == nil {
		return
	}

	// Split the rest by comma at top level, then process each sub-clause.
	clauses := splitAlterClauses(rest)

	for _, clause := range clauses {
		applyAlterClause(t, clause)
	}
}
