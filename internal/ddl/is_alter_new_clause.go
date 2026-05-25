//ff:func feature=ddl type=parse control=sequence
//ff:what ALTER TABLE 파트가 새 절의 시작인지 판별한다
package ddl

import "strings"

// isAlterNewClause checks if a part starts a new ALTER TABLE sub-clause.
func isAlterNewClause(trimmed string) bool {
	upperTrimmed := strings.ToUpper(trimmed)
	return strings.HasPrefix(upperTrimmed, "ADD COLUMN") ||
		strings.HasPrefix(upperTrimmed, "DROP COLUMN") ||
		strings.HasPrefix(upperTrimmed, "ADD CONSTRAINT") ||
		strings.HasPrefix(upperTrimmed, "DROP CONSTRAINT") ||
		strings.HasPrefix(upperTrimmed, "ALTER COLUMN") ||
		strings.HasPrefix(upperTrimmed, "RENAME ")
}
