//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 파라미터 노드에서 기본값과 기본값 호출 함수명을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractDefault extracts the default value and call function name from a parameter.
// isNone is true when the default value is Python's None literal.
func extractDefault(param *sitter.Node, src []byte) (defaultVal, defaultCall string, isNone bool) {
	for i := 0; i < int(param.ChildCount()); i++ {
		child := param.Child(i)
		defaultVal, defaultCall, isNone = tryExtractDefault(child, src)
		if defaultVal != "" || isNone {
			return
		}
	}
	return
}
