//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what 테이블의 Constraints에서 지정 컬럼의 이름 없는 FK를 제거한다
package ddl

import (
	"regexp"
	"strings"
)

// removeUnnamedFK removes unnamed FOREIGN KEY constraints that reference the
// given column from t.Constraints. A constraint is "unnamed" if it does not
// start with the CONSTRAINT keyword. This handles the case where a CREATE TABLE
// defined an inline table-level FK (e.g. "FOREIGN KEY (col) REFERENCES t(id)")
// and a later ADD CONSTRAINT replaces it with a named FK for the same column.
func removeUnnamedFK(t *Table, colName string) {
	re := regexp.MustCompile(`(?i)FOREIGN\s+KEY\s*\(\s*` + regexp.QuoteMeta(colName) + `\s*\)`)
	result := make([]string, 0, len(t.Constraints))
	for _, c := range t.Constraints {
		if re.MatchString(c) && !strings.HasPrefix(strings.TrimSpace(strings.ToUpper(c)), "CONSTRAINT") {
			continue
		}
		result = append(result, c)
	}
	t.Constraints = result
}
