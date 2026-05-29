//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 문자열 슬라이스에 값이 존재하는지 검사한다
package fastify

func containsString(s []string, v string) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}
