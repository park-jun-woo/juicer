//ff:func feature=ddl type=parse control=sequence
//ff:what 달러 태그 연속 문자 판별 (a-z, A-Z, 0-9, _)
package ddl

// isDollarTagCont returns true if ch is a valid continuation character for a dollar-quote tag identifier.
func isDollarTagCont(ch rune) bool {
	return isDollarTagStart(ch) || (ch >= '0' && ch <= '9')
}
