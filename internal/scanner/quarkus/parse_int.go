//ff:func feature=scan type=convert control=iteration dimension=1 topic=quarkus
//ff:what 문자열에서 숫자만 추출하여 int로 변환한다
package quarkus

func parseInt(s string) int {
	v := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			v = v*10 + int(c-'0')
		}
	}
	return v
}
