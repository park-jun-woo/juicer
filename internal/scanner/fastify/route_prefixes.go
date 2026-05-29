//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 라우트가 속한 wrapper 스코프 prefix를 파일 prefix와 합성하여 적용할 prefix 목록을 만든다
package fastify

func routePrefixes(r routeInfo, filePfx []string, scopes []wrapperScope) []string {
	base := filePfx
	if len(base) == 0 {
		base = []string{""}
	}
	wrapPrefix := wrapperPrefixFor(r.StartByte, scopes)
	if wrapPrefix == "" {
		return base
	}
	return composePrefixes(base, wrapPrefix)
}
