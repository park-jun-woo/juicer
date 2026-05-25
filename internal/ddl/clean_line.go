//ff:func feature=ddl type=parse control=sequence
//ff:what SQL 인라인 주석 제거 및 공백 정리
package ddl

import "strings"

// cleanLine removes inline SQL comments and trims whitespace.
func cleanLine(s string) string {
	s = strings.TrimSpace(s)
	// Remove inline comments (-- ...)
	// But be careful with strings containing --
	if idx := strings.Index(s, " --"); idx >= 0 {
		s = strings.TrimSpace(s[:idx])
	}
	return s
}
