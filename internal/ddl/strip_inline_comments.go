//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what 각 줄에서 SQL 인라인 주석 제거
package ddl

import "strings"

// stripInlineComments removes SQL inline comments (-- ...) from each line.
func stripInlineComments(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if idx := strings.Index(line, "--"); idx >= 0 {
			lines[i] = line[:idx]
		}
	}
	return strings.Join(lines, "\n")
}
