//ff:func feature=ddl type=parse control=iteration dimension=2
//ff:what 최상위 레벨에서만 구분자로 분리 (괄호 내부 무시)
package ddl

import "strings"

// splitTopLevel splits text by sep, but only at depth 0 (not inside parentheses).
func splitTopLevel(text string, sep byte) []string {
	var parts []string
	var buf strings.Builder
	depth := 0

	for i := 0; i < len(text); i++ {
		ch := text[i]
		switch ch {
		case '(':
			depth++
			buf.WriteByte(ch)
		case ')':
			depth--
			buf.WriteByte(ch)
		default:
			if ch == sep && depth == 0 {
				parts = append(parts, buf.String())
				buf.Reset()
			} else {
				buf.WriteByte(ch)
			}
		}
	}
	s := buf.String()
	if strings.TrimSpace(s) != "" {
		parts = append(parts, s)
	}
	return parts
}
