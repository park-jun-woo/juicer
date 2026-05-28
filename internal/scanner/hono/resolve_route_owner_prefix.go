//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what 라우트의 소유 변수명으로 prefix를 조회한다
package hono

func resolveRouteOwnerPrefix(r routeInfo, ctx *scanContext) string {
	if r.OwnerVar != "" {
		if p, ok := ctx.prefixMap[r.OwnerVar]; ok {
			return p
		}
	}
	return ""
}
