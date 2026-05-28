//ff:func feature=ddl type=parse control=iteration dimension=3
//ff:what 세미콜론 기준 SQL 문 분리 (괄호 깊이 + 달러 인용 추적)
package ddl

import "strings"

// splitStatements splits SQL text by semicolons, respecting parentheses
// and dollar quoting ($$...$$ and $tag$...$tag$) to avoid splitting
// inside CREATE TABLE (...) or PL/pgSQL function bodies.
func splitStatements(sql string) []string {
	var stmts []string
	var buf strings.Builder
	depth := 0
	runes := []rune(sql)
	n := len(runes)

	for i := 0; i < n; i++ {
		ch := runes[i]

		// Dollar quoting: $$ or $tag$
		if ch == '$' {
			tag := extractDollarTag(runes, i)
			if tag != "" {
				// Write the opening tag
				buf.WriteString(tag)
				i += len([]rune(tag)) - 1
				// Find the closing tag
				rest := string(runes[i+1:])
				closeIdx := strings.Index(rest, tag)
				if closeIdx >= 0 {
					buf.WriteString(rest[:closeIdx])
					buf.WriteString(tag)
					i += closeIdx + len([]rune(tag))
				}
				// If no closing tag found, just continue (unmatched dollar quote)
				continue
			}
		}

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
