//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 같은 파일에서 클래스명으로 부모를 찾아 상속 필드를 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findParentFieldsInFile looks up className in the same file and returns
// its merged (parent + own) fields.
func findParentFieldsInFile(root *sitter.Node, src []byte, className string, visited map[string]bool) []pydanticField {
	classes := findAllByType(root, "class_definition")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != className {
			continue
		}
		return resolveFieldsWithInheritance(cls, root, src, visited)
	}
	return nil
}
