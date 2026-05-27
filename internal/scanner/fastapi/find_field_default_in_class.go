//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 클래스 본문에서 필드의 기본 문자열 값을 찾는다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findFieldDefaultInClass searches a class_definition node for a field named
// attrName and returns its default string literal value.
// Supports patterns like `ATTR_NAME: type = "value"`.
func findFieldDefaultInClass(cls *sitter.Node, attrName string, src []byte) string {
	body := findChildByType(cls, "block")
	if body == nil {
		return ""
	}
	assignments := findAllByType(body, "assignment")
	for _, assign := range assignments {
		left := findChildByType(assign, "identifier")
		if left == nil || nodeText(left, src) != attrName {
			continue
		}
		strNode := findChildByType(assign, "string")
		if strNode != nil {
			return unquotePython(nodeText(strNode, src))
		}
	}
	return ""
}
