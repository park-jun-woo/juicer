//ff:func feature=scan type=test topic=flask control=iteration dimension=1
//ff:what flask 테스트 헬퍼 — 문자열 슬라이스 포함 여부
package flask

func contains(list []string, s string) bool {
	for _, e := range list {
		if e == s {
			return true
		}
	}
	return false
}
