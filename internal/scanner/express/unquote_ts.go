//ff:func feature=scan type=parse control=sequence topic=express
//ff:what TypeScript 문자열 리터럴에서 따옴표를 제거한다
package express

func unquoteTS(s string) string {
	if len(s) < 2 {
		return s
	}
	if (s[0] == '\'' && s[len(s)-1] == '\'') ||
		(s[0] == '"' && s[len(s)-1] == '"') ||
		(s[0] == '`' && s[len(s)-1] == '`') {
		return s[1 : len(s)-1]
	}
	return s
}
