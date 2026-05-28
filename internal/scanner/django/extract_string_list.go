//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what Python 리스트 노드에서 문자열 값을 추출한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractStringList extracts string values from a Python list node.
func extractStringList(listNode *sitter.Node, src []byte) []string {
	var result []string
	for i := 0; i < int(listNode.ChildCount()); i++ {
		child := listNode.Child(i)
		if child.Type() == "string" {
			val := strings.ToUpper(unquotePython(nodeText(child, src)))
			result = append(result, val)
		}
	}
	return result
}
