//ff:func feature=ddl type=parse control=sequence
//ff:what Column.Raw에서 REFERENCES table(col) 및 ON DELETE/ON UPDATE 절을 제거한다
package ddl

import (
	"regexp"
	"strings"
)

var reInlineRef = regexp.MustCompile(`(?i)\s+REFERENCES\s+\w+\s*\([^)]*\)(\s+ON\s+(?:DELETE|UPDATE)\s+(?:CASCADE|SET\s+NULL|SET\s+DEFAULT|RESTRICT|NO\s+ACTION))*`)

// stripInlineReferences removes an inline REFERENCES clause (and any trailing
// ON DELETE / ON UPDATE actions) from a column definition string.
// Example: "owner_id BIGINT NOT NULL REFERENCES owners(id) ON DELETE CASCADE"
//        → "owner_id BIGINT NOT NULL"
func stripInlineReferences(raw string) string {
	return strings.TrimSpace(reInlineRef.ReplaceAllString(raw, ""))
}
