//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what method가 "all"이면 5개 HTTP 메서드를 반환하고, 아니면 단일 메서드를 반환한다
package fastify

func expandAllMethod(method string) []string {
	if method == "all" {
		return []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	}
	return []string{method}
}
