//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 플러그인 마운트 그래프에서 파일별 transitive prefix 집합을 생성한다 (다중 등록은 별도 prefix로 유지)
package fastify

func resolvePluginPrefixes(mounts []pluginMount) map[string][]string {
	graph := buildMountGraph(mounts)
	memo := make(map[string][]string)
	for file := range graph {
		visiting := make(map[string]bool)
		filePrefixes(file, graph, memo, visiting)
	}
	return memo
}
