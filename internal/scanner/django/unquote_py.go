//ff:func feature=scan type=parse control=sequence topic=django
//ff:what Python 문자열 리터럴에서 따옴표를 제거한다
package django

import "strings"

// unquotePython strips surrounding quotes from a Python string literal.
func unquotePython(s string) string {
	if len(s) < 2 {
		return s
	}
	// Handle f-strings and raw strings
	if strings.HasPrefix(s, "f\"") || strings.HasPrefix(s, "f'") ||
		strings.HasPrefix(s, "r\"") || strings.HasPrefix(s, "r'") {
		s = s[1:]
	}
	// Triple quotes
	if len(s) >= 6 && (strings.HasPrefix(s, `"""`) || strings.HasPrefix(s, `'''`)) {
		return s[3 : len(s)-3]
	}
	if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
		return s[1 : len(s)-1]
	}
	return s
}
