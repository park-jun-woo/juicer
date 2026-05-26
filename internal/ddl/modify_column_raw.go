//ff:func feature=ddl type=parse control=selection
//ff:what Column.Raw 문자열에서 NOT NULL/DEFAULT/TYPE을 정규식으로 치환한다
package ddl

import (
	"regexp"
	"strings"
)

var (
	reNotNull       = regexp.MustCompile(`(?i)\s+NOT\s+NULL`)
	reDefaultClause = regexp.MustCompile(`(?i)\s+DEFAULT\s+(?:'[^']*'|\S+)`)
)

// modifyColumnRaw applies an ALTER COLUMN action to a column's Raw string.
// Supported actions: SET NOT NULL, DROP NOT NULL, SET DEFAULT, DROP DEFAULT, TYPE.
func modifyColumnRaw(raw, action string) string {
	upper := strings.ToUpper(strings.TrimSpace(action))

	switch {
	case upper == "SET NOT NULL":
		if reNotNull.MatchString(raw) {
			return raw
		}
		return raw + " NOT NULL"

	case upper == "DROP NOT NULL":
		return strings.TrimSpace(reNotNull.ReplaceAllString(raw, ""))

	case strings.HasPrefix(upper, "SET DEFAULT"):
		value := strings.TrimSpace(action[len("SET DEFAULT"):])
		if reDefaultClause.MatchString(raw) {
			return strings.TrimSpace(reDefaultClause.ReplaceAllString(raw, " DEFAULT "+value))
		}
		return raw + " DEFAULT " + value

	case upper == "DROP DEFAULT":
		return strings.TrimSpace(reDefaultClause.ReplaceAllString(raw, ""))

	case strings.HasPrefix(upper, "TYPE "):
		newType := strings.TrimSpace(action[len("TYPE "):])
		return changeColumnType(raw, newType)
	}

	return raw
}

