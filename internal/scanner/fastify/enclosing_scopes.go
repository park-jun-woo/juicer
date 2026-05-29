//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 주어진 바이트 위치를 포함하는 wrapper 스코프만 골라낸다
package fastify

func enclosingScopes(pos uint32, scopes []wrapperScope) []wrapperScope {
	var out []wrapperScope
	for _, ws := range scopes {
		if pos >= ws.Start && pos < ws.End {
			out = append(out, ws)
		}
	}
	return out
}
