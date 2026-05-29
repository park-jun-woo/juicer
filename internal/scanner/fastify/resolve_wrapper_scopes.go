//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 인라인 async wrapper 마운트에서 파일별 wrapper 스코프(바이트 범위+prefix) 맵을 만든다
package fastify

func resolveWrapperScopes(mounts []pluginMount) map[string][]wrapperScope {
	scopes := make(map[string][]wrapperScope)
	for _, m := range mounts {
		if !m.Inline || m.Prefix == "" || m.SourceFile == "" {
			continue
		}
		ws := wrapperScope{Start: m.WrapperStart, End: m.WrapperEnd, Prefix: m.Prefix}
		scopes[m.SourceFile] = append(scopes[m.SourceFile], ws)
	}
	return scopes
}
