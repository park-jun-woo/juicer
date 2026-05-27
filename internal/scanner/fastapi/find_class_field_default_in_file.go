//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 단일 파일에서 클래스 이름으로 필드 기본값을 찾는다
package fastapi

// findClassFieldDefaultInFile searches a single file for a class named
// className and returns the default string value for attrName.
func findClassFieldDefaultInFile(fi fileInfo, className, attrName string) string {
	classes := findAllByType(fi.root, "class_definition")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil || nodeText(nameNode, fi.src) != className {
			continue
		}
		return findFieldDefaultInClass(cls, attrName, fi.src)
	}
	return ""
}
