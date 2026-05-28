//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 모든 파일에서 ViewSet 클래스를 수집한다
package django

// collectViewSets finds all DRF ViewSet subclasses in the parsed files.
func collectViewSets(files []fileInfo) []viewsetInfo {
	var viewsets []viewsetInfo
	for _, fi := range files {
		viewsets = append(viewsets, collectViewSetsFromFile(fi)...)
	}
	return viewsets
}
