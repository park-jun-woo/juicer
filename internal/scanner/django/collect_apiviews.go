//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 모든 파일에서 APIView 클래스를 수집한다
package django

// collectAPIViews finds all DRF APIView subclasses in the parsed files. The
// class index lets transitive (custom base class) inheritance be resolved.
func collectAPIViews(files []fileInfo, idx classIndex) []apiviewInfo {
	var views []apiviewInfo
	for _, fi := range files {
		views = append(views, collectAPIViewsFromFile(fi, idx)...)
	}
	return views
}
