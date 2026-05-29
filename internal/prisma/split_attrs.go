//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what 속성 문자열을 paren/bracket 깊이 0의 @토큰 단위로 분리
package prisma

// splitAttrs splits a field attribute string into individual @attributes,
// keeping parentheses/brackets balanced inside each attribute.
func splitAttrs(s string) []string {
	attrs := make([]string, 0, 4)
	depth := 0
	start := -1
	for i := 0; i < len(s); i++ {
		c := s[i]
		depth += attrDepthDelta(c)
		if c == '@' && depth == 0 {
			attrs = appendAttr(attrs, s, start, i)
			start = i
		}
	}
	return appendAttr(attrs, s, start, len(s))
}
