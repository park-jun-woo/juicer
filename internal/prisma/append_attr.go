//ff:func feature=prisma type=parse control=sequence topic=prisma
//ff:what start이 유효한 @속성 시작이면 트림하여 attrs에 추가
package prisma

import "strings"

// appendAttr appends s[start:end] trimmed to attrs when start marks an @attribute.
func appendAttr(attrs []string, s string, start, end int) []string {
	if start < 0 {
		return attrs
	}
	if a := strings.TrimSpace(s[start:end]); a != "" {
		attrs = append(attrs, a)
	}
	return attrs
}
