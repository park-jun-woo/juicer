//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 마운트 그래프를 거슬러 올라가며 파일의 누적(transitive) prefix 집합을 계산한다
package fastify

func filePrefixes(file string, graph map[string][]pluginMount, memo map[string][]string, visiting map[string]bool) []string {
	if cached, ok := memo[file]; ok {
		return cached
	}
	if visiting[file] {
		return nil
	}
	visiting[file] = true
	var result []string
	for _, m := range graph[file] {
		parents := filePrefixes(m.SourceFile, graph, memo, visiting)
		result = appendUnique(result, composePrefixes(parents, m.Prefix)...)
	}
	delete(visiting, file)
	memo[file] = result
	return result
}
