//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what urls.py에서 urlpatterns를 수집한다
package django

// collectURLs extracts urlpatterns from all parsed files.
func collectURLs(files []fileInfo) []urlEntry {
	var allEntries []urlEntry
	for _, fi := range files {
		entries := collectURLsFromFile(fi)
		allEntries = append(allEntries, entries...)
	}
	return allEntries
}
