//ff:func feature=scan type=extract control=iteration dimension=2 topic=django
//ff:what 전 파일을 순회하여 클래스 이름→부모목록 인덱스를 구축한다
package django

// buildClassIndex walks every parsed file and records each class definition's
// direct parent class names, enabling transitive inheritance resolution.
func buildClassIndex(files []fileInfo) classIndex {
	idx := classIndex{}
	for _, fi := range files {
		for _, classNode := range findAllByType(fi.root, "class_definition") {
			nameNode := findChildByType(classNode, "identifier")
			if nameNode == nil {
				continue
			}
			name := nodeText(nameNode, fi.src)
			idx[name] = extractParentClasses(classNode, fi.src)
		}
	}
	return idx
}
