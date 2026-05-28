//ff:func feature=scan type=extract control=sequence
//ff:what isSlice이면 "[]"을 반환하고 아니면 빈 문자열을 반환한다
package echo

func slicePrefix(isSlice bool) string {
	if isSlice {
		return "[]"
	}
	return ""
}
