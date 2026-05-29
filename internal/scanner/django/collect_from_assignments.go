//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what urlpatterns = [...] 또는 i18n_patterns(...) 대입문에서 path() 호출을 수집한다
package django

// collectFromAssignments extracts urlEntries from urlpatterns = [...] assignments.
func collectFromAssignments(fi fileInfo) []urlEntry {
	var entries []urlEntry
	for _, node := range findAllByType(fi.root, "assignment") {
		if !isURLPatternsAssignment(node, fi.src) {
			continue
		}
		entries = append(entries, collectFromURLPatternsRHS(node, fi.src)...)
	}
	return entries
}
