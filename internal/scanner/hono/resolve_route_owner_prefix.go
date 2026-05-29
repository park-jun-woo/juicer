//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what 라우트의 (소유 파일, 소유 변수명)으로 prefix를 조회한다
package hono

func resolveRouteOwnerPrefix(r routeInfo, ctx *scanContext, file string) string {
	if r.OwnerVar != "" {
		if p, ok := ctx.prefixMap[prefixKey(file, r.OwnerVar)]; ok {
			return p
		}
	}
	return ""
}
