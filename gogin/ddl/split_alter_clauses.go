//ff:func feature=ddl type=parse control=iteration dimension=2
//ff:what ALTER TABLE 복합 절을 개별 절로 분리
package ddl

import "strings"

// splitAlterClauses splits ALTER TABLE sub-clauses.
// In multi-action ALTER TABLE, sub-clauses are separated by commas,
// but each new sub-clause starts with ADD or DROP keyword.
func splitAlterClauses(rest string) []string {
	// Split by comma at top level first
	parts := splitTopLevel(rest, ',')

	var clauses []string
	var currentClause strings.Builder

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if isAlterNewClause(trimmed) {
			if currentClause.Len() > 0 {
				clauses = append(clauses, currentClause.String())
				currentClause.Reset()
			}
			currentClause.WriteString(trimmed)
		} else {
			// This part is a continuation of the previous clause
			if currentClause.Len() > 0 {
				currentClause.WriteString(", ")
			}
			currentClause.WriteString(trimmed)
		}
	}
	if currentClause.Len() > 0 {
		clauses = append(clauses, currentClause.String())
	}
	return clauses
}
