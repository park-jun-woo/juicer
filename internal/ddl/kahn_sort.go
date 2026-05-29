//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 결정적 위상정렬 (준비된 노드 중 알파벳 우선 선택, 남은 순환 노드는 fallback)
package ddl

// kahnSort performs a deterministic topological sort: among nodes whose
// dependencies are already emitted, the alphabetically-first is chosen. Any
// remaining nodes form a cycle and are appended alphabetically with a warning.
func kahnSort(names []string, deps map[string]map[string]bool) []string {
	emitted := make(map[string]bool, len(names))
	result := make([]string, 0, len(names))

	for len(result) < len(names) {
		next := pickReady(names, deps, emitted)
		if next == "" {
			break
		}
		emitted[next] = true
		result = append(result, next)
	}

	return appendCycleFallback(names, emitted, result)
}
