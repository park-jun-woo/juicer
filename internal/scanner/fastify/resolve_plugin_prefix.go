//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 플러그인 마운트에서 파일별 prefix 매핑을 생성한다
package fastify

func resolvePluginPrefixes(mounts []pluginMount) map[string]string {
	prefixMap := make(map[string]string)
	for _, m := range mounts {
		if m.FilePath == "" || m.Prefix == "" {
			continue
		}
		existing := prefixMap[m.FilePath]
		if existing == "" {
			prefixMap[m.FilePath] = m.Prefix
		} else {
			prefixMap[m.FilePath] = joinFastifyPath(existing, m.Prefix)
		}
	}
	return prefixMap
}
