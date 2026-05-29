//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what Unsupported("X")의 따옴표 내부 원문을 이스케이프 해제하여 반환
package prisma

import "strings"

// unsupportedType extracts and unescapes the inner text of
// Unsupported("...") into the raw SQL type string.
func unsupportedType(t string) string {
	open := strings.IndexByte(t, '"')
	if open < 0 {
		return t
	}
	rest := t[open+1:]
	close := strings.LastIndexByte(rest, '"')
	if close < 0 {
		return t
	}
	inner := rest[:close]
	return strings.ReplaceAll(inner, `\"`, `"`)
}
