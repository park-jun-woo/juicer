//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what /* ... */ 블록 주석 제거 (여러 줄 포함)
package ddl

import "strings"

// stripBlockComments removes all /* ... */ block comments from text,
// including multi-line block comments.
func stripBlockComments(s string) string {
	var buf strings.Builder
	for {
		start := strings.Index(s, "/*")
		if start < 0 {
			buf.WriteString(s)
			break
		}
		buf.WriteString(s[:start])
		end := strings.Index(s[start+2:], "*/")
		if end < 0 {
			// Unclosed block comment — remove rest
			break
		}
		s = s[start+2+end+2:]
	}
	return buf.String()
}
