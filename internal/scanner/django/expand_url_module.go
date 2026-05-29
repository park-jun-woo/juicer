//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 한 모듈의 urlpatterns를 include를 따라 재귀 전개하여 평탄한 urlEntry로 만든다
package django

// expandURLModule recursively expands a module's urlpatterns, following include()
// and composing the prefix. visited prevents include cycles across modules.
func expandURLModule(module, prefix string, byModule map[string][]urlEntry, visited map[string]bool) []urlEntry {
	if visited[module] {
		return nil
	}
	visited[module] = true
	defer delete(visited, module)

	var out []urlEntry
	for _, entry := range byModule[module] {
		out = append(out, expandURLEntry(entry, prefix, byModule, visited)...)
	}
	return out
}
