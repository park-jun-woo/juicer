//ff:func feature=ddl type=parse control=iteration dimension=2
//ff:what 첫 번째 최상위 괄호 안의 내용 추출
package ddl

import "strings"

// extractParenBody returns the content inside the first top-level parentheses.
func extractParenBody(stmt string) string {
	start := strings.IndexByte(stmt, '(')
	if start < 0 {
		return ""
	}
	depth := 0
	for i := start; i < len(stmt); i++ {
		switch stmt[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return stmt[start+1 : i]
			}
		}
	}
	return stmt[start+1:]
}
