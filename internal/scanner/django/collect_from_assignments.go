//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what urlpatterns = [...] 대입문에서 path() 호출을 수집한다
package django

// collectFromAssignments extracts urlEntries from urlpatterns = [...] assignments.
func collectFromAssignments(fi fileInfo) []urlEntry {
	var entries []urlEntry
	for _, node := range findAllByType(fi.root, "assignment") {
		if !isURLPatternsAssignment(node, fi.src) {
			continue
		}
		listNode := findChildByType(node, "list")
		if listNode == nil {
			continue
		}
		entries = append(entries, parsePathCallsInList(listNode, fi.src)...)
	}
	return entries
}
