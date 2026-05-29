//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 슬라이스에 중복 없이 문자열들을 추가한다
package fastify

func appendUnique(dst []string, vals ...string) []string {
	for _, v := range vals {
		if !containsString(dst, v) {
			dst = append(dst, v)
		}
	}
	return dst
}
