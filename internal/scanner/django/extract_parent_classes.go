//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 클래스 정의에서 부모 클래스 이름을 추출한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractParentClasses extracts parent class names from a class definition.
func extractParentClasses(classNode *sitter.Node, src []byte) []string {
	argList := findChildByType(classNode, "argument_list")
	if argList == nil {
		return nil
	}
	var parents []string
	for i := 0; i < int(argList.ChildCount()); i++ {
		child := argList.Child(i)
		if child.Type() == "identifier" {
			parents = append(parents, nodeText(child, src))
		} else if child.Type() == "attribute" {
			text := nodeText(child, src)
			parts := strings.Split(text, ".")
			parents = append(parents, parts[len(parts)-1])
		}
	}
	return parents
}
