//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what SQL 문 앞의 주석 줄 제거 (-- 및 /* */ 블록 주석)
package ddl

import "strings"

// stripLeadingComments removes leading SQL comment lines (-- ... and /* ... */) from a statement.
func stripLeadingComments(stmt string) string {
	lines := strings.Split(stmt, "\n")
	inBlock := false
	for len(lines) > 0 {
		trimmed := strings.TrimSpace(lines[0])
		if inBlock {
			lines, inBlock = consumeBlockCommentLine(lines, trimmed)
			continue
		}
		if trimmed == "" || strings.HasPrefix(trimmed, "--") {
			lines = lines[1:]
			continue
		}
		if strings.HasPrefix(trimmed, "/*") {
			lines, inBlock = openBlockComment(lines, trimmed)
			continue
		}
		break
	}
	return strings.Join(lines, "\n")
}
