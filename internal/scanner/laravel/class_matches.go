//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what fileInfo에 주어진 이름의 클래스 선언이 있는지 보고한다
package laravel

// classMatches checks whether a fileInfo contains a class declaration matching the name.
func classMatches(fi *fileInfo, className string) bool {
	for _, cls := range findAllByType(fi.root, "class_declaration") {
		nameNode := findChildByType(cls, "name")
		if nameNode != nil && nodeText(nameNode, fi.src) == className {
			return true
		}
	}
	return false
}
