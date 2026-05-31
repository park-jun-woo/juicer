//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 모든 파일에서 ViewSet 클래스를 수집한다
package django

// collectViewSets finds all DRF ViewSet subclasses in the parsed files. The
// class index lets transitive (custom base class) inheritance be resolved.
func collectViewSets(files []fileInfo, idx classIndex) []viewsetInfo {
	var viewsets []viewsetInfo
	for _, fi := range files {
		viewsets = append(viewsets, collectViewSetsFromFile(fi, idx)...)
	}
	return viewsets
}
