//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 비인라인 마운트를 대상 파일(FilePath) 기준 인접 리스트로 구성한다
package fastify

func buildMountGraph(mounts []pluginMount) map[string][]pluginMount {
	graph := make(map[string][]pluginMount)
	for _, m := range mounts {
		if m.Inline || m.FilePath == "" {
			continue
		}
		graph[m.FilePath] = append(graph[m.FilePath], m)
	}
	return graph
}
