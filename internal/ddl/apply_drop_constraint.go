//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what DROP CONSTRAINT 절로 테이블의 Constraints에서 이름으로 제거한다
package ddl

import (
	"regexp"
	"strings"
)

// applyDropConstraint removes a named constraint from the table.
func applyDropConstraint(t *Table, constraintName string) {
	constraintName = strings.ToLower(constraintName)
	re := regexp.MustCompile(`(?i)\bCONSTRAINT\s+` + regexp.QuoteMeta(constraintName) + `\b`)
	result := make([]string, 0, len(t.Constraints))
	for _, c := range t.Constraints {
		if !re.MatchString(c) {
			result = append(result, c)
		}
	}
	t.Constraints = result
}
