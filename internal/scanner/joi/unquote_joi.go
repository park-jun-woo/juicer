//ff:func feature=scan type=parse control=sequence topic=joi
//ff:what 따옴표로 감싼 문자열 리터럴에서 따옴표를 제거한다
package joi

func unquoteJoi(s string) string {
	if len(s) < 2 {
		return s
	}
	first := s[0]
	last := s[len(s)-1]
	if (first == '\'' || first == '"' || first == '`') && first == last {
		return s[1 : len(s)-1]
	}
	return s
}
