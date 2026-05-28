//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 단일 파일에서 ViewSet 클래스를 수집한다
package django

// collectViewSetsFromFile finds ViewSet classes in a single file.
func collectViewSetsFromFile(fi fileInfo) []viewsetInfo {
	var viewsets []viewsetInfo
	for _, classNode := range findAllByType(fi.root, "class_definition") {
		vs := parseViewSetClass(classNode, fi)
		if vs != nil {
			viewsets = append(viewsets, *vs)
		}
	}
	return viewsets
}
