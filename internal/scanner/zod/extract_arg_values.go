//ff:func feature=scan type=extract control=selection topic=zod
//ff:what 단일 인자 노드에서 문자열 값을 추출한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

func extractArgValues(a *sitter.Node, src []byte) []string {
	switch a.Type() {
	case "string":
		return []string{unquoteTS(nodeText(a, src))}
	case "number":
		return []string{nodeText(a, src)}
	case "array":
		return extractArrayStringValues(a, src)
	}
	return nil
}
