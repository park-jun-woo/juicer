//ff:func feature=scan type=extract control=sequence
//ff:what method가 "any"이면 5개 HTTP 메서드 목록을, 아니면 단일 메서드 슬라이스를 반환한다
package scanner

func expandAnyMethod(method string) []string {
	if method == "any" {
		return []string{"get", "post", "put", "patch", "delete"}
	}
	return []string{method}
}
