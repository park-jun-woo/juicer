//ff:func feature=scan type=extract control=iteration dimension=1 topic=zod
//ff:what 배열 리터럴에서 문자열 요소만 추출한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

func extractArrayStringValues(a *sitter.Node, src []byte) []string {
	var result []string
	for i := 0; i < int(a.ChildCount()); i++ {
		child := a.Child(i)
		if child.Type() == "string" {
			result = append(result, unquoteTS(nodeText(child, src)))
		}
	}
	return result
}
