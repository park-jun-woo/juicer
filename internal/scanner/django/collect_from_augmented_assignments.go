//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what urlpatterns += [...] 증분 대입문에서 path() 호출을 수집한다
package django

// collectFromAugmentedAssignments extracts urlEntries from urlpatterns += [...] assignments.
func collectFromAugmentedAssignments(fi fileInfo) []urlEntry {
	var entries []urlEntry
	for _, node := range findAllByType(fi.root, "augmented_assignment") {
		leftNode := findChildByType(node, "identifier")
		if leftNode == nil || nodeText(leftNode, fi.src) != "urlpatterns" {
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
