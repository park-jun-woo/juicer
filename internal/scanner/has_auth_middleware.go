//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what Middleware 목록에서 인증 키워드가 포함된 항목이 있는지 확인한다
package scanner

func hasAuthMiddleware(middlewares []string) bool {
	for _, mw := range middlewares {
		if containsAuthKeyword(mw) {
			return true
		}
	}
	return false
}
