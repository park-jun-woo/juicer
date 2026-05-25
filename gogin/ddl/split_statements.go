//ff:func feature=ddl type=parse control=iteration dimension=3
//ff:what 세미콜론 기준 SQL 문 분리 (괄호 깊이 추적)
package ddl

import "strings"

// splitStatements splits SQL text by semicolons, respecting parentheses
// to avoid splitting inside CREATE TABLE (...);
func splitStatements(sql string) []string {
	var stmts []string
	var buf strings.Builder
	depth := 0

	for _, ch := range sql {
		switch ch {
		case '(':
			depth++
			buf.WriteRune(ch)
		case ')':
			depth--
			buf.WriteRune(ch)
		case ';':
			if depth <= 0 {
				s := strings.TrimSpace(buf.String())
				if s != "" {
					stmts = append(stmts, s)
				}
				buf.Reset()
				depth = 0
			} else {
				buf.WriteRune(ch)
			}
		default:
			buf.WriteRune(ch)
		}
	}
	s := strings.TrimSpace(buf.String())
	if s != "" {
		stmts = append(stmts, s)
	}
	return stmts
}
