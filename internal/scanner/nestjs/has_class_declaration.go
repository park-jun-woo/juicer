//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 파일에 해당 이름의 class 선언이 있는지 확인한다
package nestjs

import "os"

// hasClassDeclaration reports whether filePath contains a class declaration
// with the given className.
func hasClassDeclaration(filePath, className string) bool {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}
	root, err := parseTypeScript(src)
	if err != nil {
		return false
	}
	for _, cls := range findAllByType(root, "class_declaration") {
		nameNode := findChildByType(cls, "type_identifier")
		if nameNode != nil && nodeText(nameNode, src) == className {
			return true
		}
	}
	return false
}
