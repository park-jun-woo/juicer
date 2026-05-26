//ff:func feature=ddl type=extract control=sequence
//ff:what FK 정의에서 컬럼명을 추출한다 (FOREIGN KEY (col) → col)
package ddl

import "regexp"

var reFKColumn = regexp.MustCompile(`(?i)FOREIGN\s+KEY\s*\(\s*(\w+)\s*\)`)

// extractFKColumn extracts the column name from a FOREIGN KEY constraint definition.
// Given "CONSTRAINT x FOREIGN KEY (owner_id) REFERENCES users(id)", it returns "owner_id".
// Returns "" if no FOREIGN KEY clause is found.
func extractFKColumn(constraintDef string) string {
	m := reFKColumn.FindStringSubmatch(constraintDef)
	if m == nil {
		return ""
	}
	return m[1]
}
