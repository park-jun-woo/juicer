//ff:func feature=scan type=extract control=sequence
//ff:what Endpoint가 인증 필요 엔드포인트인지 판별한다
package scanner

func isAuthEndpoint(ep Endpoint) bool {
	if ep.AuthLevel == "auth_required" {
		return true
	}
	if len(ep.Roles) > 0 {
		return true
	}
	if ep.AuthLevel == "public" {
		return false
	}
	return hasAuthMiddleware(ep.Middleware)
}
