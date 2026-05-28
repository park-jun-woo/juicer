//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what Pass 1: 파일 파싱, Fastify 인스턴스 수집, 플러그인 마운트 수집, import 해석
package fastify

func scanPass1(tsFiles []string, absRoot string) *scanContext {
	parsed := make(map[string]*fileInfo)
	allInstances := make(map[string]map[string]bool)
	var allMounts []pluginMount
	for _, path := range tsFiles {
		r := scanOneFilePass1(path, absRoot)
		if r == nil {
			continue
		}
		parsed[path] = r.fi
		allInstances[path] = r.instances
		allMounts = append(allMounts, r.mounts...)
	}
	prefixMap := resolvePluginPrefixes(allMounts)
	return &scanContext{parsed: parsed, instances: allInstances, prefixMap: prefixMap, absRoot: absRoot}
}
