//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what list 노드에서 따옴표 제거한 문자열 값을 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// extractStringLiterals extracts unquoted string values from a list node.
func extractStringLiterals(listNode *sitter.Node, src []byte) []string {
	var result []string
	for i := 0; i < int(listNode.ChildCount()); i++ {
		child := listNode.Child(i)
		if child.Type() == "string" {
			result = append(result, unquotePython(nodeText(child, src)))
		}
	}
	return result
}
