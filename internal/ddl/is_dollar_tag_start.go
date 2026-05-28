//ff:func feature=ddl type=parse control=sequence
//ff:what 달러 태그 시작 문자 판별 (a-z, A-Z, _)
package ddl

// isDollarTagStart returns true if ch is a valid first character for a dollar-quote tag identifier.
func isDollarTagStart(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}
