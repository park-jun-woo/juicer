//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what SQL 문 앞의 주석 줄 제거
package ddl

import "strings"

// stripLeadingComments removes leading SQL comment lines (-- ...) from a statement.
func stripLeadingComments(stmt string) string {
	lines := strings.Split(stmt, "\n")
	for len(lines) > 0 {
		trimmed := strings.TrimSpace(lines[0])
		if trimmed == "" || strings.HasPrefix(trimmed, "--") {
			lines = lines[1:]
		} else {
			break
		}
	}
	return strings.Join(lines, "\n")
}
