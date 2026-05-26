//ff:func feature=scan type=parse control=sequence topic=fastapi
//ff:what 문자가 타입 분할자(콤마 또는 파이프)인지 확인한다
package fastapi

// isSplitter returns true if the character is a comma or pipe.
func isSplitter(ch rune) bool {
	return ch == ',' || ch == '|'
}
