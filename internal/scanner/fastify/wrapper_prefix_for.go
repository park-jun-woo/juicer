//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 주어진 바이트 위치를 감싸는 wrapper 스코프들의 prefix를 바깥→안쪽 순으로 합성한다
package fastify

func wrapperPrefixFor(pos uint32, scopes []wrapperScope) string {
	enclosing := enclosingScopes(pos, scopes)
	sortScopesByWidth(enclosing)
	prefix := ""
	for _, ws := range enclosing {
		prefix = joinFastifyPath(prefix, ws.Prefix)
	}
	return prefix
}
