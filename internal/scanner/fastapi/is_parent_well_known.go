//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 같은 파일 내 부모 클래스가 well-known을 상속하는지 확인한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// isParentWellKnown checks whether parentName is defined in the same file
// and inherits from a well-known base class.
func isParentWellKnown(root *sitter.Node, src []byte, parentName string) bool {
	classes := findAllByType(root, "class_definition")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != parentName {
			continue
		}
		return hasWellKnownParent(cls, src)
	}
	return false
}
