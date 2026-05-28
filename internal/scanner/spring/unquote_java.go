//ff:func feature=scan type=convert control=sequence topic=spring
//ff:what Java 문자열 리터럴의 따옴표를 제거한다
package spring

func unquoteJava(s string) string {
	if len(s) < 2 {
		return s
	}
	if s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
