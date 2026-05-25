//ff:func feature=sql type=parse control=sequence
//ff:what 연속 공백을 단일 공백으로 정규화
package sqls

import (
	"strings"
)

// normalizeWhitespace collapses runs of whitespace into single spaces.
//
func normalizeWhitespace(s string) string {
	fields := strings.Fields(s)
	return strings.Join(fields, " ")
}

