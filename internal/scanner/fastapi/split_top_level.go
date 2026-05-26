//ff:func feature=scan type=parse control=iteration dimension=1 topic=fastapi
//ff:what 타입 문자열을 최상위 레벨에서 콤마나 파이프로 분할한다
package fastapi

import "strings"

// splitTopLevel splits a type string by commas or pipes at the top level
// (not inside brackets).
func splitTopLevel(s string) []string {
	var parts []string
	depth := 0
	start := 0
	for i, ch := range s {
		depth += bracketDelta(ch)
		if isSplitter(ch) && depth == 0 {
			parts = append(parts, strings.TrimSpace(s[start:i]))
			start = i + 1
		}
	}
	tail := strings.TrimSpace(s[start:])
	if tail != "" {
		parts = append(parts, tail)
	}
	return parts
}
