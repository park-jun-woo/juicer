//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 단일 파일에서 APIView 클래스를 수집한다
package django

// collectAPIViewsFromFile finds APIView classes in a single file.
func collectAPIViewsFromFile(fi fileInfo) []apiviewInfo {
	var views []apiviewInfo
	for _, classNode := range findAllByType(fi.root, "class_definition") {
		av := parseAPIViewClass(classNode, fi)
		if av != nil {
			views = append(views, *av)
		}
	}
	return views
}
