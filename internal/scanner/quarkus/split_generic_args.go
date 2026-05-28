//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 제네릭 인자 문자열을 콤마 기준으로 분리한다
package quarkus

import "strings"

func splitGenericArgs(s string) []string {
	var parts []string
	depth := 0
	start := 0
	for i, ch := range s {
		depth += angleBracketDelta(ch)
		if ch == ',' && depth == 0 {
			parts = append(parts, strings.TrimSpace(s[start:i]))
			start = i + 1
		}
	}
	parts = append(parts, strings.TrimSpace(s[start:]))
	return parts
}
